package service

import "loan-gateway/gateway/internal/client"

type Service interface {
	Check(info LoanInfo) (*client.InferenceResponse, error)
}
