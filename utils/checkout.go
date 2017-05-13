package utils

// CheckOut interface sets discount strategy and total signature
type CheckOut interface {
	SetDiscountStrategy(DiscountStrategy)
	Total() (string, error)
}

// DiscountStrategy struct
type DiscountStrategy func(string, []string) (string, error)

// order struct to hold checkout information
type order struct {
	customerID string
	productIDs []string
	total      string
	strategy   DiscountStrategy
	errorCode  error
}

// SetDiscountStrategy sets discount strategy for checkOut
func (o *order) SetDiscountStrategy(s DiscountStrategy) {
	o.strategy = s
}

// Total calculates total for checkOut
func (o *order) Total() (string, error) {
	o.total, o.errorCode = o.strategy(o.customerID, o.productIDs)
	return o.total, o.errorCode
}

// NewCheckOut returns a CheckOut interface
func NewCheckOut(customerID string, products []string) CheckOut {
	return &order{customerID: customerID, productIDs: products}
}
