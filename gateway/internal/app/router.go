package app

import "net/http"

func (a *App) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /",
		ChainMiddleware(a.HandleIndex,
			RecoveryMiddleware,
			LoggingMiddleware,
			RateLimitMiddleware))

	mux.HandleFunc("POST /loan-check",
		ChainMiddleware(a.HandleLoanCheck,
			RecoveryMiddleware,
			LoggingMiddleware,
			RateLimitMiddleware))
	return mux
}
