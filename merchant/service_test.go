package merchant

import (
	"context"
	"testing"
)

func TestInquiry(t *testing.T) {
	srv, ctx := setup()
	inq, err := srv.Inquiry(ctx)
	if err != nil {
		t.Errorf("Error:%s", err)
	}
	ok := inq == "inquiry"
	if !ok {
		t.Errorf("expected service to be inquiry")
	}
}

func TestPayment(t *testing.T) {
	srv, ctx := setup()
	pay, err := srv.Payment(ctx)
	if err != nil {
		t.Errorf("Error:%s", err)
	}
	ok := pay == "payment"
	if !ok {
		t.Errorf("expected service to be inquiry")
	}
}
func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
