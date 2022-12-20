package handlers

import "net/http"

//IPaymentIntent ...
type IPaymentIntent interface {
	StartPaymentIntent(w http.ResponseWriter, r *http.Request)
}
