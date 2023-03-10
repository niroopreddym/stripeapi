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
	r.Handle("/customers", http.HandlerFunc(customerHandler.AddCustomer)).Methods("POST")
	r.Handle("/customers", http.HandlerFunc(customerHandler.ListCustomers)).Methods("GET")
	r.Handle("/customers/{customer_id}", http.HandlerFunc(customerHandler.DeleteCustomer)).Methods("GET")
	r.Handle("/customers/{customer_id}", http.HandlerFunc(customerHandler.PutCustomer)).Methods("PUT")
	r.Handle("/customers/{customer_id}", http.HandlerFunc(customerHandler.DeleteCustomer)).Methods("DELETE")

	paymentIntentHandler := handlers.NewPaymentIntentHandler()
	r.Handle("/paymentintent", http.HandlerFunc(paymentIntentHandler.StartPaymentIntent)).Methods("POST")

	SessionHandler := handlers.NewSesssionsHandler()
	r.Handle("/session", http.HandlerFunc(SessionHandler.CreateSession)).Methods("POST")

	log.Fatal(http.ListenAndServe(":9294", r))
}
