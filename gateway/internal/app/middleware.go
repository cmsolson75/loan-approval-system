package app

import (
	"loan-gateway/internal/config"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Middleware defines a function that wraps an http.HandlerFunc.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// LoggingMiddleware logs the duration of each HTTP request.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// RecoveryMiddleware recovers from panics and returns 500 Internal Server Error.
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic: %v", rec)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu       sync.Mutex
	visitors = make(map[string]*visitor)
)

// StartRateLimiterCleanup launches a goroutine to periodically clean up expired rate limiters.
func StartRateLimiterCleanup(cfg *config.Config) {
	go func() {
		ticker := time.NewTicker(cfg.RateCleanup)
		defer ticker.Stop()

		for range ticker.C {
			mu.Lock()
			for ip, v := range visitors {
				if time.Since(v.lastSeen) > cfg.RateTTL {
					delete(visitors, ip)
				}
			}
			mu.Unlock()
		}
	}()
}

func getVisitor(cfg *config.Config, ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if v, ok := visitors[ip]; ok {
		v.lastSeen = time.Now()
		return v.limiter
	}
	limiter := rate.NewLimiter(rate.Every(cfg.RateInterval), cfg.RateBurst)
	visitors[ip] = &visitor{
		limiter:  limiter,
		lastSeen: time.Now(),
	}
	return limiter
}

// RateLimitMiddleware returns middleware that applies IP-based rate limiting.
func RateLimitMiddleware(cfg *config.Config) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort((r.RemoteAddr))
			if err != nil {
				http.Error(w, "unable to determine client IP", http.StatusInternalServerError)
				return
			}
			limiter := getVisitor(cfg, ip)
			if !limiter.Allow() {
				log.Printf("BLOCKED: rate limit exceeded for IP %s", ip)
				http.Error(w, "Too many Requests", http.StatusTooManyRequests)
				return
			}
			log.Printf("ALLOWED: %s", ip)

			next(w, r)
		}
	}

}

// ChainMiddleware applies a sequence of middlewares to a handler.
func ChainMiddleware(h http.HandlerFunc, mws ...Middleware) http.HandlerFunc {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
