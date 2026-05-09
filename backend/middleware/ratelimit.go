package middleware

import (
	"net"
	"net/http"
	"resu-me/service"
	"strings"
)

func RateLimit(rl *service.RateLimiter, trustedProxies []*net.IPNet) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := extractIP(r, trustedProxies)
			allowed, _ := rl.Allow(ip)
			if !allowed {
				w.Header().Set("Retry-After", "60")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func extractIP(r *http.Request, trustedProxies []*net.IPNet) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	if len(trustedProxies) > 0 {
		remoteIP := net.ParseIP(host)
		if remoteIP != nil {
			for _, cidr := range trustedProxies {
				if cidr.Contains(remoteIP) {
					if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
						parts := strings.Split(fwd, ",")
						return strings.TrimSpace(parts[0])
					}
					break
				}
			}
		}
	}

	return host
}
