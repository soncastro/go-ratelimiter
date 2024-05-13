package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Middleware struct {
	limiter *Limiter
	keyType string
}

func (m *Middleware) LimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		if m.keyType == "TokenRateLimiter" {
			key = c.GetHeader("Authorization")
		}

		allowed, err := m.limiter.CheckLimit(key, m.keyType)
		if err != nil || !allowed {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "você atingiu o número máximo de solicitações ou ações permitidas dentro de um determinado período de tempo",
			})
			return
		}

		c.Next()
	}
}

func NewRedisDatastore() Datastore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &RedisDatastore{
		client: rdb,
	}
}

func main() {
	rateLimitPerSec, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		rateLimitPerSec = 1
	}

	rateLimiter := NewLimiter(NewRedisDatastore(), rateLimitPerSec)
	rateLimiterMiddleware := &Middleware{
		limiter: rateLimiter,
		keyType: "RateLimiter",
	}

	tokenRateLimitPerSec, err := strconv.Atoi(os.Getenv("TOKEN_RATE_LIMIT"))
	if err != nil {
		tokenRateLimitPerSec = 1
	}

	tokenRateLimiter := NewLimiter(NewRedisDatastore(), tokenRateLimitPerSec)
	tokenRateLimiterMiddleware := &Middleware{
		limiter: tokenRateLimiter,
		keyType: "TokenRateLimiter",
	}

	r := gin.Default()

	r.Use(rateLimiterMiddleware.LimiterMiddleware())
	r.Use(tokenRateLimiterMiddleware.LimiterMiddleware())

	r.GET("/example-endpoint", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success!",
		})
	})

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
