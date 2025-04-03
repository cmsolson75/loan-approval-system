package service

import (
	"fmt"
	"loan-gateway/gateway/internal/client"
	"strconv"
)

type LoanInfo struct {
	AnnualIncome string
	LoanAmount   string
	LoanTerm     string
	CreditScore  string
}

type LoanService struct {
	Client *client.InferenceClient
}

func NewLoanService(c *client.InferenceClient) *LoanService {
	return &LoanService{Client: c}
}

func (s *LoanService) Check(info LoanInfo) (*client.InferenceResponse, error) {
	income, err := validateNonNegativeField("Annual Income", info.AnnualIncome)
	if err != nil {
		return nil, err
	}

	amount, err := validateNonNegativeField("Loan Amount", info.LoanAmount)
	if err != nil {
		return nil, err
	}
	term, err := validateIntField("Loan Term", info.LoanTerm, 2, 30)
	if err != nil {
		return nil, err
	}
	score, err := validateIntField("Credit Score", info.CreditScore, 300, 900)
	if err != nil {
		return nil, err
	}
	input := client.InferenceRequest{
		AnnualIncome: income,
		LoanAmount:   amount,
		LoanTerm:     term,
		CreditScore:  score,
	}

	return s.Client.Predict(input)
}

func validateIntField(name, value string, min, max int) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid integer", name)
	}
	if v < min || v > max {
		return 0, fmt.Errorf("%s must be between %d and %d", name, min, max)
	}
	return v, nil
}

func validateNonNegativeField(name, value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid integer", name)
	}
	if v < 0 {
		return 0, fmt.Errorf("%s can't be negative", name)
	}
	return v, nil
}
