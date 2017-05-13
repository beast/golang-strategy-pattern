package utils

import (
	"strategy-pattern/model"
	"strings"

	"github.com/shopspring/decimal"
)

// GetDiscountStrategy get discount strategy by customer id.
func GetDiscountStrategy(customerID string) DiscountStrategy {
	discountRuleID := model.GetDiscountRule(customerID)
	switch discountRuleID {
	// In production environment the rule id should be uuid, using string switch here to simulate a rule engine
	case "1":
		return xForY(3, "classic")
	case "2":
		return priceDrop("299.99", "standout", 0)
	case "3":
		return priceDrop("379.99", "premium", 4)
	case "4":
		break
	default:
		return noDiscount()
	}
	return nil
}

// decorators
// no discounts
func noDiscount() DiscountStrategy {
	return func(customerID string, products []string) (string, error) {
		var total decimal.Decimal
		for _, productID := range products {
			// todo: productID validation
			price, err := model.GetProductPrice(productID)
			if err != nil {
				return "", err
			}
			total = total.Add(price)
		}
		return total.StringFixed(2), nil
	}
}

// x for x-1 discounts
func xForY(x int, productType string) DiscountStrategy {
	return func(customerID string, productIDs []string) (string, error) {
		var total decimal.Decimal
		var productCounter int
		for _, productID := range productIDs {
			if strings.EqualFold(productType, productID) {
				productCounter++
			}
			// todo: productID validation
			price, err := model.GetProductPrice(productID)
			if err != nil {
				return "", err
			}
			if productCounter%x != 0 || !strings.EqualFold(productType, productID) {
				total = total.Add(price)
			}
		}
		return total.StringFixed(2), nil
	}
}

// price drop discounts
func priceDrop(newPrice string, productType string, dropStartsAt int) DiscountStrategy {
	return func(customerID string, productIDs []string) (string, error) {
		var total decimal.Decimal
		var price decimal.Decimal
		var err error
		productMap := productCount(productIDs)
		for _, productID := range productIDs {
			// todo: productID validation
			price, err = model.GetProductPrice(productID)
			if err != nil {
				return "", err
			}
			if strings.EqualFold(productType, productID) && productMap[productType] >= dropStartsAt {
				price, err = decimal.NewFromString(newPrice)
				if err != nil {
					return "", err
				}
			}
			total = total.Add(price)
		}
		return total.StringFixed(2), nil
	}
}

// Counts the number of products in same category and put in a map
func productCount(list []string) map[string]int {

	duplicateFrequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicateFrequency[item]

		if exist {
			duplicateFrequency[item]++ // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}
