package utils

import (
	"strategy-pattern/model"

	"github.com/shopspring/decimal"
)

// GetDiscountStrategy get discount strategy by customer id.
func GetDiscountStrategy(customerID string) DiscountStrategy {
	discountRuleID := model.GetDiscountRule(customerID)
	switch discountRuleID {
	// In production environment the rule id should be uuid, using string switch here to simulate a rule engine
	case "1":
		return unileverDiscount()
	case "2":
		return appleDiscount()
	case "3":
		return nikeDiscount()
	case "4":
		return fordDiscount()
	default:
		return defaultDiscount()
	}
	return nil
}

// decorators
func NoDiscount(count int, price decimal.Decimal) decimal.Decimal {
	return price.Mul(decimal.NewFromFloat(float64(count)))
}

func XForY(count int, x int, y int, price decimal.Decimal) decimal.Decimal {
	return price.Mul(decimal.NewFromFloat(float64(count/x*y + count%x)))
}

func PriceDrop(count int, oldPrice decimal.Decimal, sNewPrice string, dropStartsAt int) decimal.Decimal {
	newPrice, _ := decimal.NewFromString(sNewPrice)
	if count >= dropStartsAt {
		return newPrice.Mul(decimal.NewFromFloat(float64(count)))
	}
	return oldPrice.Mul(decimal.NewFromFloat(float64(count)))
}

// Counts the number of products in same category and put in a map
func getProductCount(list []string) map[string]int {

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

// todo: during production the followings should read the rules from database
// discount strategies
func defaultDiscount() DiscountStrategy {
	return func(customerID string, productIDs []string) (string, error) {
		var total string
		productCountMap := getProductCount(productIDs)
		// todo: should put prices in a map instead of getting it one by one
		classicPrice, err := model.GetProductPrice("classic")
		if err != nil {
			return "", err
		}
		standoutPrice, err := model.GetProductPrice("standout")
		if err != nil {
			return "", err
		}
		premiumPrice, err := model.GetProductPrice("premium")
		if err != nil {
			return "", err
		}
		classicTotal := NoDiscount(productCountMap["classic"], classicPrice)
		standoutTotal := NoDiscount(productCountMap["standout"], standoutPrice)
		premiumTotal := NoDiscount(productCountMap["premium"], premiumPrice)
		total = classicTotal.Add(standoutTotal).Add(premiumTotal).StringFixed(2)
		return total, nil
	}
}

func unileverDiscount() DiscountStrategy {
	return func(customerID string, productIDs []string) (string, error) {
		var total string
		productCountMap := getProductCount(productIDs)
		// todo: should put prices in a map instead of getting it one by one
		classicPrice, err := model.GetProductPrice("classic")
		if err != nil {
			return "", err
		}
		standoutPrice, err := model.GetProductPrice("standout")
		if err != nil {
			return "", err
		}
		premiumPrice, err := model.GetProductPrice("premium")
		if err != nil {
			return "", err
		}
		classicTotal := XForY(productCountMap["classic"], 3, 2, classicPrice)
		standoutTotal := NoDiscount(productCountMap["standout"], standoutPrice)
		premiumTotal := NoDiscount(productCountMap["premium"], premiumPrice)
		total = classicTotal.Add(standoutTotal).Add(premiumTotal).StringFixed(2)
		return total, nil
	}
}

func appleDiscount() DiscountStrategy {
	return func(customerID string, productIDs []string) (string, error) {
		var total string
		productCountMap := getProductCount(productIDs)
		// todo: should put prices in a map instead of getting it one by one
		classicPrice, err := model.GetProductPrice("classic")
		if err != nil {
			return "", err
		}
		standoutPrice, err := model.GetProductPrice("standout")
		if err != nil {
			return "", err
		}
		premiumPrice, err := model.GetProductPrice("premium")
		if err != nil {
			return "", err
		}
		classicTotal := NoDiscount(productCountMap["classic"], classicPrice)
		standoutTotal := PriceDrop(productCountMap["standout"], standoutPrice, "299.99", 0)
		premiumTotal := NoDiscount(productCountMap["premium"], premiumPrice)
		total = classicTotal.Add(standoutTotal).Add(premiumTotal).StringFixed(2)
		return total, nil
	}
}

func nikeDiscount() DiscountStrategy {
	return func(customerID string, productIDs []string) (string, error) {
		var total string
		productCountMap := getProductCount(productIDs)
		// todo: should put prices in a map instead of getting it one by one
		classicPrice, err := model.GetProductPrice("classic")
		if err != nil {
			return "", err
		}
		standoutPrice, err := model.GetProductPrice("standout")
		if err != nil {
			return "", err
		}
		premiumPrice, err := model.GetProductPrice("premium")
		if err != nil {
			return "", err
		}
		classicTotal := NoDiscount(productCountMap["classic"], classicPrice)
		standoutTotal := NoDiscount(productCountMap["standout"], standoutPrice)
		premiumTotal := PriceDrop(productCountMap["premium"], premiumPrice, "379.99", 4)
		total = classicTotal.Add(standoutTotal).Add(premiumTotal).StringFixed(2)
		return total, nil
	}
}

func fordDiscount() DiscountStrategy {
	return func(customerID string, productIDs []string) (string, error) {
		var total string
		productCountMap := getProductCount(productIDs)
		// todo: should put prices in a map instead of getting it one by one
		classicPrice, err := model.GetProductPrice("classic")
		if err != nil {
			return "", err
		}
		standoutPrice, err := model.GetProductPrice("standout")
		if err != nil {
			return "", err
		}
		premiumPrice, err := model.GetProductPrice("premium")
		if err != nil {
			return "", err
		}
		classicTotal := XForY(productCountMap["classic"], 5, 4, classicPrice)
		standoutTotal := PriceDrop(productCountMap["standout"], standoutPrice, "309.99", 0)
		premiumTotal := PriceDrop(productCountMap["premium"], premiumPrice, "389.99", 3)
		total = classicTotal.Add(standoutTotal).Add(premiumTotal).StringFixed(2)
		return total, nil
	}
}
