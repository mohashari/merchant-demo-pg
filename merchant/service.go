package merchant

import (
	"context"
)

type Service interface {
	Inquiry(ctx context.Context) (string, error)
	Payment(ctx context.Context) (string, error)
}

type transactionService struct{}

func NewService() Service {
	return transactionService{}
}

func (transactionService) Inquiry(ctx context.Context) (string, error) {
	return "inquiry", nil
}

func (transactionService) Payment(ctx context.Context) (string, error) {
	return "payment", nil
}
