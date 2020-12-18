package merchant

import (
	"context"
	"encoding/json"
	"net/http"
)

type transactionResponse struct {
	Status string `json:"status"`
	Err    string `json:"err,omitempty"`
}

type transactionRequest struct{}

func decodeTransactionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transactionRequest
	return req, nil
}

func encodeTransactionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)

}
