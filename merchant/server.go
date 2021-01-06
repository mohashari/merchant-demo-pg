package merchant

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, endpoint Endpoints) http.Handler {

	r := mux.NewRouter()
	r.Use(commondMidleware)

	r.Methods("GET").Path("/inquiry").Handler(httptransport.NewServer(
		endpoint.Inquiry,
		decodeTransactionRequest,
		encodeTransactionResponse,
	))

	r.Methods("GET").Path("/payment").Handler(httptransport.NewServer(
		endpoint.Payment,
		decodeTransactionRequest,
		encodeTransactionResponse,
	))

	return r
}

func commondMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application-json")
		next.ServeHTTP(w, r)
	})
}
