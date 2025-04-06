package app

import (
	"html/template"
	"loan-gateway/internal/service"
)

// App holds shared application state including templates and loan service.
type App struct {
	Templates   *template.Template
	LoanService service.Service
}

// New returns a new App instance with parsed templates and provided loan service.
func New(svc service.Service) *App {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	return &App{
		Templates:   tmpl,
		LoanService: svc,
	}
}
