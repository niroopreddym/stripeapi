package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niroopreddym/StripeAPI/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	customerHandler := handlers.NewStripeCustomerHandler()
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	r.Handle("/customers/addcustomer", http.HandlerFunc(customerHandler.AddCustomer)).Methods("POST")
	r.Handle("/customers", http.HandlerFunc(customerHandler.ListCustomers)).Methods("GET")
	r.Handle("/customers/updatecustomer", http.HandlerFunc(customerHandler.PutCustomer)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":9294", r))
}