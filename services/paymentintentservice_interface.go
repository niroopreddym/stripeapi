package services

import (
	"github.com/niroopreddym/StripeAPI/models"
	"github.com/stripe/stripe-go"
)

// IPaymentIntentServices ....
type IPaymentIntentServices interface {
	CreateNewPaymentIntent(payment models.PaymentIntent) (*stripe.PaymentIntent, error)
	ConfirmPaymentIntent(payment models.PaymentIntent) (*stripe.PaymentIntent, error)
}
