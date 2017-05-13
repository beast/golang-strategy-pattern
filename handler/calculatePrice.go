package handler

import (
	"net/http"
	"strategy-pattern/utils"

	"github.com/labstack/echo"
)

// CalculatePrice base on customer id and products purchased.
func CalculatePrice(c echo.Context) error {
	var i struct {
		// Shall use uuid in the production
		CustomerID string   `json:"customerID" valid:"alpha,required"`
		Products   []string `json:"products" valid:"alpha,required"`
	}
	err := Bind(c, &i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	checkOut := utils.NewCheckOut(i.CustomerID, i.Products)
	discountStrategy := utils.GetDiscountStrategy(i.CustomerID)
	checkOut.SetDiscountStrategy(discountStrategy)
	total, err := checkOut.Total()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": total,
	})
}
