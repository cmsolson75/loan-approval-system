package main

import (
	"loan-gateway/gateway/internal/app"
	"log"
	"net/http"
)

func main() {
	app.StartRateLimiterCleanup()
	gateway := app.New("http://127.0.0.1:8000/predict")
	log.Println("Gateway listening on :8020")
	log.Fatal(http.ListenAndServe(":8020", gateway.Router()))
}
