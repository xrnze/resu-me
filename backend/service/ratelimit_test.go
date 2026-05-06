package service

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimiter_NewIPGetsFreshBucket(t *testing.T) {
	rl := NewRateLimiter(10, 3, 0)
	rl.GetBucket("1.2.3.4")
	allowed, _ := rl.Allow("1.2.3.4")
	assert.True(t, allowed, "new IP should get a fresh bucket")
}

func TestRateLimiter_AllowsWithinLimit(t *testing.T) {
	rl := NewRateLimiter(100, 5, 0)
	for i := 0; i < 5; i++ {
		allowed, _ := rl.Allow("10.0.0.1")
		assert.True(t, allowed, "request %d should be allowed within burst", i)
	}
}

func TestRateLimiter_BlocksOverLimit(t *testing.T) {
	rl := NewRateLimiter(100, 2, 0)
	rl.Allow("10.0.0.2")
	rl.Allow("10.0.0.2")
	allowed, _ := rl.Allow("10.0.0.2")
	assert.False(t, allowed, "request over limit should be blocked")
}

func TestRateLimiter_DifferentIPsIndependent(t *testing.T) {
	rl := NewRateLimiter(100, 2, 0)
	rl.Allow("10.0.0.1")
	rl.Allow("10.0.0.1")
	allowed, _ := rl.Allow("10.0.0.2")
	assert.True(t, allowed, "different IP should not be affected by another IP's limit")
}

func TestRateLimiter_Cleanup(t *testing.T) {
	rl := NewRateLimiter(100, 2, 50*time.Millisecond)
	rl.Allow("10.0.0.1")
	rl.Allow("10.0.0.2")
	time.Sleep(100 * time.Millisecond)
	rl.Cleanup()
	rl.mu.Lock()
	count := len(rl.limiters)
	rl.mu.Unlock()
	assert.Equal(t, 0, count, "expected 0 buckets after cleanup")
}

func TestRateLimiter_ConcurrentAccess(t *testing.T) {
	rl := NewRateLimiter(100, 10, 0)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rl.Allow("10.0.0.1")
		}(i)
	}
	wg.Wait()
}
