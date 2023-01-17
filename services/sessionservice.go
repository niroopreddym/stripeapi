package services

import (
	"github.com/niroopreddym/StripeAPI/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"github.com/stripe/stripe-go/client"
)

// StripeSessionService ...
type StripeSessionService struct {
	StripeAPI *client.API
}

// NewStripeSessionService ...
func NewStripeSessionService(stripeAPI *client.API) ISessionService {
	return &StripeSessionService{
		StripeAPI: stripeAPI,
	}
}

// CreateNewSession ...
func (ss *StripeSessionService) CreateNewSession(stripeSession models.StripeSession) (*stripe.CheckoutSession, error) {
	checkoutSessionParams := stripe.CheckoutSessionParams{
		CancelURL:          &stripeSession.CancelURL,
		Locale:             stripe.String("en"),
		SuccessURL:         &stripeSession.SuccessURL,
		Customer:           &stripeSession.CustomerID,
		PaymentMethodTypes: stripe.StringSlice(stripeSession.PaymentMethodTypes),
		Mode:               &stripeSession.Mode,
	}

	lstSesionLineItems := []*stripe.CheckoutSessionLineItemParams{}

	for _, lineItem := range stripeSession.LineItems {
		checkoutLineItem := stripe.CheckoutSessionLineItemParams{
			Quantity: stripe.Int64(int64(lineItem.Quantity)),
		}

		lstSesionLineItems = append(lstSesionLineItems, &checkoutLineItem)
	}

	checkoutSessionParams.LineItems = lstSesionLineItems
	return session.New(&checkoutSessionParams)
}
