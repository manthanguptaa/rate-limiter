package main

import (
	"sync"
	"time"
)

type TokenBucket struct {
	tokens         int
	maxTokens      int
	refillRate     int
	lastRefillTime time.Time
	mutex          sync.RWMutex
}

const (
	maxTokens  = 10
	refillRate = 1
)

func InitializeTokenBucket() *TokenBucket {
	return &TokenBucket{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) Refill() {
	now := time.Now()
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	if (tb.tokens < tb.maxTokens) && (time.Since(tb.lastRefillTime) >= time.Second) {
		tb.tokens = tb.tokens + tb.refillRate
		tb.lastRefillTime = now
	}
}

func (tb *TokenBucket) IsRequestAllowed() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}
