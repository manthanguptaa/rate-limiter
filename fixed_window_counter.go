package main

import (
	"sync"
	"time"
)

type FixedWindowCounter struct {
	windowStart  time.Time
	windowSize   int
	currentCount int
	maxRequest   int
	mu           sync.Mutex
}

const (
	windowSize = 60
	maxRequest = 60
)

func InitialiseFixedWindowCounter() *FixedWindowCounter {
	return &FixedWindowCounter{
		windowStart:  time.Now().Truncate(time.Second * 60),
		windowSize:   windowSize,
		currentCount: 0,
		maxRequest:   maxRequest,
	}
}

func (fwc *FixedWindowCounter) IsRequestAllowed() bool {
	fwc.mu.Lock()
	defer fwc.mu.Unlock()

	now := time.Now()
	window_start := now.Truncate(time.Second * 60)
	isInWindow := now.Sub(fwc.windowStart) < time.Duration(fwc.windowSize)*time.Second
	if isInWindow {
		if fwc.currentCount < fwc.maxRequest {
			fwc.currentCount++
			return true
		} else {
			return false
		}
	} else {
		fwc.windowStart = window_start
		fwc.currentCount = 1
		return true
	}
}
