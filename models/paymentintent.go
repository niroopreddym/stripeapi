package models

import "github.com/stripe/stripe-go"

//PaymentIntent ...
type PaymentIntent struct {
	PaymentID     string
	Amount        int64
	Currency      stripe.Currency
	PaymentMethod string
}
