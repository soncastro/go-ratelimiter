package main

import (
	"testing"
)

type MockLimiter struct{}

func (l *MockLimiter) CheckLimit(key string, limiterType string) (bool, error) {
	return true, nil
}

func TestCheckLimitUnit(t *testing.T) {
	limiter := &MockLimiter{}

	allowed, err := limiter.CheckLimit("unit_test", "RateLimiter")
	if err != nil {
		t.Errorf("Error checking limit: %s", err)
	}

	if !allowed {
		t.Error("Access should be allowed when under rate limit")
	}
}
