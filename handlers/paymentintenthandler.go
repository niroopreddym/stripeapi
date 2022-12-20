package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/niroopreddym/StripeAPI/clients"
	"github.com/niroopreddym/StripeAPI/models"
	"github.com/niroopreddym/StripeAPI/services"
)

// PaymentIntentHandler ...
type PaymentIntentHandler struct {
	PaymentIntentService services.IPaymentIntentServices
}

// NewPaymentIntentHandler is the ctor
func NewPaymentIntentHandler() IPaymentIntent {
	clientAPIInstance := clients.NewStripeClient()
	return &PaymentIntentHandler{
		PaymentIntentService: services.NewPaymentIntentService(clientAPIInstance),
	}
}

// StartPaymentIntent start a payment processing request
func (h *PaymentIntentHandler) StartPaymentIntent(w http.ResponseWriter, r *http.Request) {
	customerPayload := models.PaymentIntent{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		responseController(w, http.StatusInternalServerError, readErr)
		return
	}

	strBufferValue := string(bodyBytes)
	err := json.Unmarshal([]byte(strBufferValue), &customerPayload)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	responsePayload, err := h.PaymentIntentService.CreateNewPaymentIntent(customerPayload)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, map[string]string{
		"paymentIntentID": *&responsePayload.ID,
	})
}
