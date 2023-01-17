package models

// StripeSession ...
type StripeSession struct {
	CancelURL          string     `json:"cancel_url"`
	SuccessURL         string     `json:"success_url"`
	CustomerID         string     `json:"customer"`
	LineItems          []LineItem `json:"lineitems"`
	Mode               string     `json:"mode"`
	PaymentMethodTypes []string   `json:"payment_method_types"` //can be enum of card and pm_card_visa
	IdempotencyKey     string     `json:"Idempotency-Key"`
}

//LineItem ..
type LineItem struct {
	PriceData []PriceData `json:"price_data"`
	Quantity  int         `json:"quantity"`
}

// PriceData ...
type PriceData struct {
	Currency    string      `json:"currency"`
	ProductData ProductData `json:"product_data"`
	UnitAmount  int         `json:"unit_amount"`
}

//ProductData ...
type ProductData struct {
	Name string `json:"name"`
}

// PaymentIntentData ...
type PaymentIntentData struct {
	ReceiptEmail string `json:"receipt_email"`
}
