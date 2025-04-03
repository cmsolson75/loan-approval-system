package app

import (
	"loan-gateway/gateway/internal/config"
	"net/http"
)

func (a *App) Router(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()
	rateLimitMw := RateLimitMiddleware(cfg)
	mux.HandleFunc("GET /",
		ChainMiddleware(a.HandleIndex,
			RecoveryMiddleware,
			LoggingMiddleware,
			rateLimitMw))

	mux.HandleFunc("POST /loan-check",
		ChainMiddleware(a.HandleLoanCheck,
			RecoveryMiddleware,
			LoggingMiddleware,
			rateLimitMw))
	return mux
}
