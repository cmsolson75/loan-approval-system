package main

import (
	"fmt"
	"loan-gateway/gateway/internal/app"
	"loan-gateway/gateway/internal/client"
	"loan-gateway/gateway/internal/config"
	"loan-gateway/gateway/internal/service"
	"log"
	"net/http"
)

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
