package model

import (
	"errors"

	"github.com/shopspring/decimal"
)

// GetProductPrice simulate a function call to get product price from Database table
func GetProductPrice(productID string) (decimal.Decimal, error) {
	priceTable := map[string]string{
		"classic":  "269.99",
		"standout": "322.99",
		"premium":  "394.99",
	}
	if value, ok := priceTable[productID]; ok {
		price, err := decimal.NewFromString(value)
		if err != nil {
			return decimal.Zero, err
		}
		return price, nil
	}
	return decimal.Zero, errors.New("Could not find the product: " + productID)
}

// GetDiscountRule simulate a function call to get discount rule id from Database table
func GetDiscountRule(customerID string) (string, error) {
	// todo: This is only simulating a customer/discount type mapping. In implementation this should be using a database and uuid
	customerDiscountTable := map[string]string{
		"Unilever": "1", // Gets a for 3 for 2 deal on Classic Ads
		"Apple":    "2", // Gets a discount on Standout Ads where the price drops to $299.99 per ad
		"Nike":     "3", // Gets a discount on Premium Ads where 4 or more are purchased. The price drops to $379.99 per ad
		"Ford":     "4", // TLDR
	}
	if value, ok := customerDiscountTable[customerID]; ok {
		return value, nil
	}
	return "", errors.New("Could not find customer: " + customerID)
}
