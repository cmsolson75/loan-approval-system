package app

import (
	"html/template"
	"loan-gateway/gateway/internal/client"
	"loan-gateway/gateway/internal/service"
)

type App struct {
	Templates   *template.Template
	LoanService *service.LoanService
}

func New(apiURL string) *App {
	tmpl := template.Must(template.ParseGlob("internal/templates/*.html"))
	client := client.NewClient(apiURL)
	service := service.NewLoanService(client)
	return &App{
		Templates:   tmpl,
		LoanService: service,
	}
}
