package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type InferenceRequest struct {
	AnnualIncome int `json:"annual_income"`
	LoanAmount   int `json:"loan_amount"`
	LoanTerm     int `json:"loan_term"`
	CreditScore  int `json:"credit_score"`
}

type InferenceResponse struct {
	ApprovalStatus string  `json:"approval_status"`
	Confidence     float64 `json:"confidence"`
}

type InferenceClient struct {
	Endpoint string
	HTTP     *http.Client
}

func NewClient(endpoint string) *InferenceClient {
	return &InferenceClient{
		Endpoint: endpoint,
		HTTP:     &http.Client{},
	}
}

func (c *InferenceClient) Predict(input InferenceRequest) (*InferenceResponse, error) {
	body, _ := json.Marshal(input)
	resp, err := c.HTTP.Post(c.Endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr struct {
			Detail []struct {
				Msg string `json:"msg"`
			} `json:"detail"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return nil, errors.New("inference API returned error")
		}

		if len(apiErr.Detail) > 0 {
			return nil, errors.New(apiErr.Detail[0].Msg)
		}
		return nil, errors.New("inference API returned an unknown validation error")
	}

	var result InferenceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
