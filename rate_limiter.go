package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type LimiterInfo struct {
	Count       int
	LastRefresh time.Time
}

type Datastore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type RedisDatastore struct {
	client *redis.Client
}

func (r *RedisDatastore) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RedisDatastore) Set(key string, value string) error {
	return r.client.Set(key, value, 0).Err()
}

type Limiter struct {
	datastore Datastore
	ratelimit int
}

func NewLimiter(datastore Datastore, ratelimit int) *Limiter {
	return &Limiter{
		datastore: datastore,
		ratelimit: ratelimit,
	}
}

func (l *Limiter) CheckLimit(key string, limitType string) (bool, error) {

	infoStr, err := l.datastore.Get(key)
	if err != nil {
		panic(err)
	}

	var info LimiterInfo
	if infoStr == "" {
		info = LimiterInfo{
			Count:       1,
			LastRefresh: time.Now(),
		}
	} else {
		err = json.Unmarshal([]byte(infoStr), &info)
		if err != nil {
			return false, err
		}
	}

	if time.Now().Sub(info.LastRefresh) > time.Minute {
		info.Count = 1
		info.LastRefresh = time.Now()
	}

	// Additional logic:
	if info.Count > l.ratelimit {
		now := time.Now()

		// Do we have a block in place?
		blockedUntilStr, err := l.datastore.Get("BLOCKED_" + key)
		if err == nil {
			blockedUntil, err := time.Parse(time.RFC3339Nano, blockedUntilStr)
			if err == nil {
				// There is a block in place.
				if now.Before(blockedUntil) {
					return false, fmt.Errorf("rate limit exceeded, too many requests")
				}

				// Block has expired - clear it.
				_ = l.datastore.Set("BLOCKED_"+key, "")
			}
		}

		// Set new block of 5 mins
		blockedUntil := now.Add(5 * time.Minute)
		_ = l.datastore.Set("BLOCKED_"+key, blockedUntil.Format(time.RFC3339Nano))

		return false, fmt.Errorf("rate limit exceeded, too many requests")
	} else {
		info.Count++
	}

	marshalInfo, _ := json.Marshal(info)
	infoStr = string(marshalInfo)
	err = l.datastore.Set(key, infoStr)
	if err != nil {
		return false, err
	}

	return true, nil
}
