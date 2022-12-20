package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/niroopreddym/StripeAPI/clients"
	"github.com/niroopreddym/StripeAPI/models"
	"github.com/niroopreddym/StripeAPI/services"
)

// StripeCustomerHandler ...
type StripeCustomerHandler struct {
	customerService services.ICustomerService
}

// NewStripeCustomerHandler  is the ctor
func NewStripeCustomerHandler() ICustomerHandler {
	clientAPIInstance := clients.NewStripeClient()
	return &StripeCustomerHandler{
		customerService: services.NewCustomerService(clientAPIInstance),
	}
}

// AddCustomer adds customoe to the stripe management account
func (h *StripeCustomerHandler) AddCustomer(w http.ResponseWriter, r *http.Request) {
	//from mux extract the payload details
	customerModel := models.Customer{}
	customerDetails, err := h.customerService.AddCustomer(customerModel)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, map[string]string{
		"customerID": *&customerDetails.ID,
	})
}

// ListCustomers ...
func (h *StripeCustomerHandler) ListCustomers(w http.ResponseWriter, r *http.Request) {
	lstCustomers, err := h.customerService.ListCustomers()
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, lstCustomers)
}

// PutCustomer ...
func (h *StripeCustomerHandler) PutCustomer(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}
	_, err := h.customerService.UpdateCustomer(customer)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusNoContent, "Updated Sucessfully")
}

func responseController(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
