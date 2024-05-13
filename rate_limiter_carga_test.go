package main

import (
	"encoding/json"
	"testing"
)

type LimiterInfoTest struct {
	Limit int `json:"limit"`
	Count int `json:"count"`
}

type MockDatastore struct {
	data map[string]*LimiterInfoTest
}

func NewMockDatastore() *MockDatastore {
	return &MockDatastore{
		data: make(map[string]*LimiterInfoTest),
	}
}

func (m *MockDatastore) Get(key string) (string, error) {
	limiterInfo := m.data[key]
	if limiterInfo == nil {
		limiterInfo = &LimiterInfoTest{Limit: 0, Count: 0}
	}
	bytes, err := json.Marshal(limiterInfo)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (m *MockDatastore) Set(key string, value string) error {
	limiterInfo := &LimiterInfoTest{}
	err := json.Unmarshal([]byte(value), limiterInfo)
	if err != nil {
		return err
	}
	m.data[key] = limiterInfo
	return nil
}

func TestCheckLimit(t *testing.T) {
	mockDatastore := NewMockDatastore()
	rateLimiter := NewLimiter(mockDatastore, 5)
	allowed, err := rateLimiter.CheckLimit("test_key", "RateLimiter")

	if err != nil {
		t.Errorf("Erro ao verificar o limite: %s", err)
	}

	if !allowed {
		t.Error("Deve ser permitido quando abaixo do limite de taxa")
	}

	for i := 0; i < 5; i++ {
		allowed, _ = rateLimiter.CheckLimit("test_key", "RateLimiter")
	}
}
