package services

import (
	"github.com/niroopreddym/StripeAPI/models"
	"github.com/stripe/stripe-go"
)

// ICustomerService ....
type ICustomerService interface {
	AddCustomer(customer models.Customer) (*stripe.Customer, error)
	ListCustomers() ([]*stripe.Customer, error)
	UpdateCustomer(customer models.Customer) (*stripe.Customer, error)
	DeleteCustomer(customerID string) error
	GetCustomerByID(customerID string) (*stripe.Customer, error)
}
