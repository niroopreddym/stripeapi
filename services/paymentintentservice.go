package services

import (
	"github.com/niroopreddym/StripeAPI/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
	"github.com/stripe/stripe-go/paymentintent"
)

// PaymentIntentService ...
type PaymentIntentService struct {
	StripeAPI *client.API
}

// NewPaymentIntentService ...
func NewPaymentIntentService(stripeAPI *client.API) IPaymentIntentServices {
	return &PaymentIntentService{
		StripeAPI: stripeAPI,
	}
}

// CreateNewPaymentIntent ...
func (h *PaymentIntentService) CreateNewPaymentIntent(payment models.PaymentIntent) (*stripe.PaymentIntent, error) {
	params := stripe.PaymentIntentParams{
		Amount:   stripe.Int64(payment.Amount),
		Currency: stripe.String(string(payment.Currency)),
	}

	return paymentintent.New(&params)
}
