package app

import "net/http"

func (a *App) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", WithLogging(a.HandleIndex))
	// API Name improvement
	mux.HandleFunc("POST /loan-check", WithLogging(a.HandleLoanCheck))
	return mux
}
