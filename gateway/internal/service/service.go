package service

import "loan-gateway/internal/client"

// Service defines the interface for checking loan approval.
type Service interface {
	Check(info LoanInfo) (*client.InferenceResponse, error)
}
