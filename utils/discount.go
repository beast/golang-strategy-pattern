package utils

import (
	"log"
	"strategy-pattern/model"

	"github.com/shopspring/decimal"
)

// GetDiscountStrategy get discount strategy by customer id.
func GetDiscountStrategy(customerID string) DiscountStrategy {
	discountRuleID := model.GetDiscountRule(customerID)
	log.Println("discount rule: " + discountRuleID)
	switch discountRuleID {
	// In production environment the rule id should be uuid, using string switch here to simulate a rule engine
	case "1":
		break
	case "2":
		break
	case "3":
		break
	case "4":
		break
	default:
		log.Println("default strategy")
		return noDiscount()
	}
	return nil
}

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

func xForY(x int, y int) DiscountStrategy {
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
