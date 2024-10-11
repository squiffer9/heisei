package middleware

import (
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type RateLimiterMiddleware struct {
	logger   *zap.Logger
	visitors map[string]*rate.Limiter
	mu       sync.RWMutex
	r        rate.Limit
	b        int
}

func NewRateLimiterMiddleware(r rate.Limit, b int, logger *zap.Logger) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		logger:   logger,
		visitors: make(map[string]*rate.Limiter),
		r:        r,
		b:        b,
	}
}

func (m *RateLimiterMiddleware) getVisitor(ip string) *rate.Limiter {
	m.mu.Lock()
	defer m.mu.Unlock()

	limiter, exists := m.visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(m.r, m.b)
		m.visitors[ip] = limiter
	}

	return limiter
}

func (m *RateLimiterMiddleware) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		limiter := m.getVisitor(ip)
		if !limiter.Allow() {
			m.logger.Warn("Rate limit exceeded",
				zap.String("ip", ip),
				zap.String("path", r.URL.Path),
			)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *RateLimiterMiddleware) CleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		m.mu.Lock()
		for ip, limiter := range m.visitors {
			if time.Since(limiter.LastEvent()) > time.Hour {
				delete(m.visitors, ip)
			}
		}
		m.mu.Unlock()
	}
}
