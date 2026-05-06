package service

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type limiterEntry struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type RateLimiter struct {
	mu       sync.Mutex
	limiters map[string]*limiterEntry
	rate     float64
	burst    int
	ttl      time.Duration
	stopCh   chan struct{}
}

func NewRateLimiter(r float64, burst int, ttl time.Duration) *RateLimiter {
	rl := &RateLimiter{
		limiters: make(map[string]*limiterEntry),
		rate:     r,
		burst:    burst,
		ttl:      ttl,
		stopCh:   make(chan struct{}),
	}
	if ttl > 0 {
		go rl.cleanupLoop()
	}
	return rl
}

func (rl *RateLimiter) GetBucket(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	entry, exists := rl.limiters[ip]
	if !exists {
		entry = &limiterEntry{
			limiter:  rate.NewLimiter(rate.Limit(rl.rate), rl.burst),
			lastSeen: time.Now(),
		}
		rl.limiters[ip] = entry
	}
	entry.lastSeen = time.Now()
	return entry.limiter
}

func (rl *RateLimiter) Allow(ip string) (bool, time.Duration) {
	limiter := rl.GetBucket(ip)
	return limiter.Allow(), 0
}

func (rl *RateLimiter) Cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	for ip, entry := range rl.limiters {
		if now.Sub(entry.lastSeen) > rl.ttl {
			delete(rl.limiters, ip)
		}
	}
}

func (rl *RateLimiter) Stop() {
	close(rl.stopCh)
}

func (rl *RateLimiter) cleanupLoop() {
	ticker := time.NewTicker(rl.ttl / 2)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			rl.Cleanup()
		case <-rl.stopCh:
			return
		}
	}
}
