package service

import (
	"errors"
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
	income, err := strconv.Atoi(info.AnnualIncome)
	if err != nil || income <= 0 {
		return nil, errors.New("invalid income")
	}
	amount, err := strconv.Atoi(info.LoanAmount)
	if err != nil || amount <= 0 {
		return nil, errors.New("invalid loan amount")
	}
	term, err := strconv.Atoi(info.LoanTerm)
	if err != nil || term < 2 || term > 30 {
		return nil, errors.New("invalid loan term")
	}
	score, err := strconv.Atoi(info.CreditScore)
	if err != nil || score < 300 || score > 900 {
		return nil, errors.New("invalid credit score")
	}
	input := client.InferenceRequest{
		AnnualIncome: income,
		LoanAmount:   amount,
		LoanTerm:     term,
		CreditScore:  score,
	}

	return s.Client.Predict(input)
}
