package middleware

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"resu-me/service"
)

func TestRateLimit_Allowed(t *testing.T) {
	rl := service.NewRateLimiter(100, 3, 0)
	handler := RateLimit(rl, nil)(testHandler())

	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "10.0.0.1:12345"
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRateLimit_Blocked(t *testing.T) {
	rl := service.NewRateLimiter(100, 1, 0)
	handler := RateLimit(rl, nil)(testHandler())

	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "10.0.0.2:12345"

	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req)
	assert.Equal(t, http.StatusOK, w1.Code)

	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req)
	assert.Equal(t, http.StatusTooManyRequests, w2.Code)
}

func TestRateLimit_RetryAfterHeader(t *testing.T) {
	rl := service.NewRateLimiter(100, 1, 0)
	handler := RateLimit(rl, nil)(testHandler())

	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "10.0.0.3:12345"

	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req)

	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req)

	assert.NotEmpty(t, w2.Header().Get("Retry-After"), "expected Retry-After header on 429")
}

func TestRateLimit_DifferentIPs(t *testing.T) {
	rl := service.NewRateLimiter(100, 1, 0)
	handler := RateLimit(rl, nil)(testHandler())

	req1 := httptest.NewRequest("GET", "/", nil)
	req1.RemoteAddr = "10.0.0.4:12345"
	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)

	req2 := httptest.NewRequest("GET", "/", nil)
	req2.RemoteAddr = "10.0.0.5:12345"
	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusOK, w2.Code, "different IP got status %d, want 200", w2.Code)
}

func TestRateLimit_ForwardedForIgnoredWithoutTrust(t *testing.T) {
	rl := service.NewRateLimiter(100, 1, 0)
	handler := RateLimit(rl, nil)(testHandler())

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "203.0.113.1")
	req.RemoteAddr = "10.0.0.1:12345"

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRateLimit_ForwardedForTrustedProxy(t *testing.T) {
	_, trustedNet, _ := net.ParseCIDR("10.0.0.0/8")
	rl := service.NewRateLimiter(100, 1, 0)
	handler := RateLimit(rl, []*net.IPNet{trustedNet})(testHandler())

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "203.0.113.1")
	req.RemoteAddr = "10.0.0.1:12345"

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
