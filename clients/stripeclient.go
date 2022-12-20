package clients

import (
	"os"

	"github.com/stripe/stripe-go/client"
)

// NewStripeClient is ctor for the client intialisation
func NewStripeClient() *client.API {
	privateKey := os.Getenv("sk_key")
	sc := client.API{}
	sc.Init(privateKey, nil)
	return &sc
}
