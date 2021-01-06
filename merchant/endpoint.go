package merchant

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Inquiry endpoint.Endpoint
	Payment endpoint.Endpoint
}

func MakeEndpoint(s Service) Endpoints {
	return Endpoints{
		Inquiry: MakeInquiryEndpoint(s),
		Payment: MakePaymentEndpoint(s),
	}
}

func MakeInquiryEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transactionRequest)
		i, err := srv.Inquiry(ctx)
		if err != nil {
			return transactionResponse{i, err.Error()}, nil
		}
		return transactionResponse{i, ""}, nil
	}
}

func MakePaymentEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transactionRequest)
		p, err := srv.Payment(ctx)
		if err != nil {
			return transactionResponse{p, err.Error()}, nil
		}
		return transactionResponse{p, ""}, nil
	}
}

// func (e Endpoints) Inquiry(ctx context.Context) (string, error) {
// 	req := transactionRequest{}
// 	resp, err := e.Inquiry(ctx, req)
// 	if err != nil {
// 		return "", err
// 	}
// 	getResp := resp.(transactionResponse)
// 	if getResp.Err != "" {
// 		return "", err
// 	}
// 	return getResp.Status, nil
// }
// func (e Endpoints) Payment(ctx context.Context) (string, error) {
// 	req := transactionRequest{}
// 	resp, err := e.Payment(ctx, req)
// 	if err != nil {
// 		return "", err
// 	}
// 	getResp := resp.(transactionResponse)
// 	if getResp.Err != "" {
// 		return getResp.Status.nil
// 	}
// }
