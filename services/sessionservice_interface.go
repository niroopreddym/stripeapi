package services

import (
	"github.com/niroopreddym/StripeAPI/models"
	"github.com/stripe/stripe-go"
)

// ISessionService ...
type ISessionService interface {
	CreateNewSession(stripeSession models.StripeSession) (*stripe.CheckoutSession, error)
}
