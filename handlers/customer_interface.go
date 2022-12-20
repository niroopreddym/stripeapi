package handlers

import (
	"net/http"
)

// ICustomerHandler handler interface
type ICustomerHandler interface {
	AddCustomer(w http.ResponseWriter, r *http.Request)
	ListCustomers(w http.ResponseWriter, r *http.Request)
	PutCustomer(w http.ResponseWriter, r *http.Request)
}
