package main

import (
	"fmt"
	"loan-gateway/internal/app"
	"loan-gateway/internal/client"
	"loan-gateway/internal/config"
	"loan-gateway/internal/service"
	"log"
	"net/http"
)

// main initializes configuration, services, and starts the HTTP server.
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}
	app.StartRateLimiterCleanup(cfg)

	var inferenceClient client.Client = client.NewClient(cfg.APIEndpoint)
	var loanService service.Service = service.NewLoanService(inferenceClient)

	gateway := app.New(loanService)
	port := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Gateway listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, gateway.Router(cfg)))
}
