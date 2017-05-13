package handler

import (
	"net/http"

	"strings"

	"github.com/labstack/echo"
	"github.com/shopspring/decimal"
)

// CalculatePrice base on customer id and products purchased.
func CalculatePrice(c echo.Context) error {
	var i struct {
		CustomerID string `json:"customerID" valid:"uuid,required"`
		// Shall update the product table to use uuid in the coding practice
		Products []string `json:"products" valid:"alpha,required"`
	}
	err := Bind(c, &i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	checkOut := NewCheckOut(i.CustomerID, i.Products)
	strategy, err := GetDiscountStrategyByID(i.CustomerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	checkOut.SetDiscountStrategy(strategy)
	total := checkOut.Total()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": total,
	})
}

// GetDiscountStrategyByID get discount strategy by customer id.
func GetDiscountStrategyByID(customerID string) DiscountStrategy {
	discountRuleID := 
	switch {
	case strings.EqualFold(, "1"):
	case strings.EqualFold(customerDiscountTable[customerID], "2"):
	case strings.EqualFold(customerDiscountTable[customerID], "3"):
	case strings.EqualFold(customerDiscountTable[customerID], "4"):
	default:
		return func(customerID string, products []string) (string, error) {
			var total string
			for _, productID := range products {
				price, err := getProductPrice(productID)
				if err != nil {
					return "", err
				}
			}
			return total, nil
		}
	}
	return nil
}

// simulate a product table
func getProductPrice(productID string) (decimal.Decimal, error) {
	priceTable := map[string]string{
		"classic":  "269.99",
		"standout": "322.99",
		"premium":  "394.99",
	}
	price, err := decimal.NewFromString(priceTable[productID])
	if err != nil {
		return decimal.Zero, err
	}
	return price, nil
}

// simulate customer's discount rule table, returns discount rule id
func getDiscountRule(customerID string) string {
	// todo: This is only simulating a customer/discount type mapping. In implementation this should be using a database and uuid
	customerDiscountTable := map[string]string{
		"Unilever": "1", // Gets a for 3 for 2 deal on Classic Ads
		"Apple":    "2", // Gets a discount on Standout Ads where the price drops to $299.99 per ad
		"Nike":     "3", // Gets a discount on Premium Ads where 4 or more are purchased. The price drops to $379.99 per ad
		"Ford":     "4", // TLDR
	}
	return customerDiscountTable[customerID]
}

// NewCheckOut returns a CheckOut interface
func NewCheckOut(customerID string, products []string) CheckOut {
	return &checkOut{customerID: customerID, products: products}
}

// DiscountStrategy struct
type DiscountStrategy func(string, []string) (string, error)

// CheckOut interface sets discount strategy and total signature
type CheckOut interface {
	SetDiscountStrategy(DiscountStrategy)
	Total() string
}

// checkOut concrete struct
type checkOut struct {
	customerID string
	products   []string
	total      string
	strategy   DiscountStrategy
}

// SetDiscountStrategy sets discount strategy for checkOut
func (ps *checkOut) SetDiscountStrategy(s DiscountStrategy) {
	ps.strategy = s
}

// Total calculates total for checkOut
func (ps *checkOut) Total() string {
	ps.total = ps.strategy(ps.customerID, ps.products)
	return ps.total
}
