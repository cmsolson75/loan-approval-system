package app

import (
	"log"
	"net/http"
)

// Render executes a named template and writes it to the response.
func (a *App) Render(w http.ResponseWriter, name string, data any) {
	err := a.Templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Printf("Template Error: %v", err)
		http.Error(w, "Template render error", http.StatusInternalServerError)
	}
}
