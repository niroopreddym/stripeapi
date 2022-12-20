package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	customerPayload := models.Customer{}
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

	responsePayload, err := h.customerService.AddCustomer(customerPayload)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, map[string]string{
		"customerID": *&responsePayload.ID,
	})
}

// GetCustomerByID gets a customer by ID
func (h *StripeCustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID, isExtsts := params["customer_id"]
	if !isExtsts {
		log.Println("missing customer ID")
		responseController(w, http.StatusBadRequest, "missing customer ID")
		return
	}

	customerDetails, err := h.customerService.GetCustomerByID(customerID)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	//you can wrap the response json to the required JSON Output
	responseController(w, http.StatusOK, customerDetails)
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

	params := mux.Vars(r)
	customerID, isExtsts := params["customer_id"]
	if !isExtsts {
		log.Println("missing customer ID")
		responseController(w, http.StatusBadRequest, "missing customer ID")
		return
	}

	customerPayload := models.Customer{}
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

	customerPayload.ID = customerID
	_, err = h.customerService.UpdateCustomer(customerPayload)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusNoContent, "Updated Sucessfully")
}

// DeleteCustomer ...
func (h *StripeCustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	customerID, isExtsts := params["id"]
	if !isExtsts {
		log.Println("missing customer ID")
		responseController(w, http.StatusBadRequest, "missing customer ID")
		return
	}

	err := h.customerService.DeleteCustomer(customerID)
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
