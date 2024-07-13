package stripe

import (
	"os"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
)

// StripeClient interface defines the methods our client will implement
type StripeClient interface {
	CreatePaymentIntent(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
	RetrievePaymentIntent(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
	ConfirmPaymentIntent(id string, params *stripe.PaymentIntentConfirmParams) (*stripe.PaymentIntent, error)
	CapturePaymentIntent(id string, params *stripe.PaymentIntentCaptureParams) (*stripe.PaymentIntent, error)
	CancelPaymentIntent(id string, params *stripe.PaymentIntentCancelParams) (*stripe.PaymentIntent, error)
}

// stripeClient struct implements the StripeClient interface
type stripeClient struct{}

// NewStripeClient returns a new instance of stripeClient
func NewStripeClient() StripeClient {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	return &stripeClient{}
}

func (s *stripeClient) CreatePaymentIntent(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return paymentintent.New(params)
}

func (s *stripeClient) RetrievePaymentIntent(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return paymentintent.Get(id, params)
}

func (s *stripeClient) ConfirmPaymentIntent(id string, params *stripe.PaymentIntentConfirmParams) (*stripe.PaymentIntent, error) {
	return paymentintent.Confirm(id, params)
}

func (s *stripeClient) CapturePaymentIntent(id string, params *stripe.PaymentIntentCaptureParams) (*stripe.PaymentIntent, error) {
	return paymentintent.Capture(id, params)
}

func (s *stripeClient) CancelPaymentIntent(id string, params *stripe.PaymentIntentCancelParams) (*stripe.PaymentIntent, error) {
	return paymentintent.Cancel(id, params)
}
