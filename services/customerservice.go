package services

import (
	"github.com/niroopreddym/StripeAPI/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

// CustomerService ...
type CustomerService struct {
	StripeAPI *client.API
}

// NewCustomerService ...
func NewCustomerService(stripeAPI *client.API) ICustomerService {
	return &CustomerService{
		StripeAPI: stripeAPI,
	}
}

// AddCustomer adds customoe to the stripe management account
func (h *CustomerService) AddCustomer(customer models.Customer) (*stripe.Customer, error) {
	customerParams := stripe.CustomerParams{
		Name:  &customer.Name,
		Email: &customer.Email,
	}

	sCustomer, err := h.StripeAPI.Customers.New(&customerParams)
	if err != nil {
		return nil, err
	}

	return sCustomer, nil
}

// ListCustomers lists all the existing customers in the system
func (h *CustomerService) ListCustomers() ([]*stripe.Customer, error) {
	params := stripe.CustomerListParams{}
	params.Single = true
	i := h.StripeAPI.Customers.List(&params)
	lstCustomers := []*stripe.Customer{}
	for i.Next() {
		c := i.Customer()
		lstCustomers = append(lstCustomers, c)
	}

	return lstCustomers, nil
}

// UpdateCustomer ...
func (h *CustomerService) UpdateCustomer(customer models.Customer) (*stripe.Customer, error) {
	params := stripe.CustomerParams{
		Email: &customer.Email,
	}

	cust, err := h.StripeAPI.Customers.Update(customer.ID, &params)
	if err != nil {
		return nil, err
	}

	return cust, nil
}

// DeleteCustomer ...
func (h *CustomerService) DeleteCustomer(customerID string) error {
	_, err := h.StripeAPI.Customers.Del(customerID, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetCustomerByID ...
func (h *CustomerService) GetCustomerByID(customerID string) (*stripe.Customer, error) {
	return h.StripeAPI.Customers.Get(customerID, nil)
}
