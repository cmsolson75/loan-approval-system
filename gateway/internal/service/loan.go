package service

import (
	"fmt"
	"loan-gateway/internal/client"
	"strconv"
)

// LoanInfo holds user-submitted loan application data.
type LoanInfo struct {
	AnnualIncome string
	LoanAmount   string
	LoanTerm     string
	CreditScore  string
}

// LoanService provides loan checking using a prediction client.
type LoanService struct {
	Client client.Client
}

var _ Service = (*LoanService)(nil)

// NewLoanService returns a new LoanService using the provided client.
func NewLoanService(c client.Client) *LoanService {
	return &LoanService{Client: c}
}

// Check validates loan input and invokes the client to get a prediction.
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
