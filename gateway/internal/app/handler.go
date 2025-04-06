package app

import (
	"fmt"
	"loan-gateway/internal/service"
	"net/http"
)

// HandleIndex renders the index page.
func (a *App) HandleIndex(w http.ResponseWriter, r *http.Request) {
	a.Render(w, "index.html", nil)
}

// HandleLoanCheck processes loan form input and renders prediction result.
func (a *App) HandleLoanCheck(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	info := service.LoanInfo{
		AnnualIncome: r.FormValue("income"),
		LoanAmount:   r.FormValue("amount"),
		LoanTerm:     r.FormValue("term"),
		CreditScore:  r.FormValue("score"),
	}
	// bad name: need to refactor
	result, err := a.LoanService.Check(info)
	if err != nil {
		a.Render(w, "text-display.html", map[string]any{"Message": err.Error()})
		return
	}
	msg := fmt.Sprintf("Status %s | Confidence: %.4f", result.ApprovalStatus, result.Confidence)
	a.Render(w, "text-display.html", map[string]any{"Message": msg})
}
