package models

import "github.com/stripe/stripe-go"

//PaymentIntent ...
type PaymentIntent struct {
	Amount   int64
	Currency stripe.Currency
}
