package middleware

import (
	"go.uber.org/zap"
	"net/http"
)

type Middleware struct {
	Logger *zap.Logger
}

func NewMiddleware(logger *zap.Logger) *Middleware {
	return &Middleware{
		Logger: logger,
	}
}

func (m *Middleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Logger.Info("Received request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)
		next.ServeHTTP(w, r)
	})
}

