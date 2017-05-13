package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

// Bind API parameters to context, validates data type according to json tag
func Bind(c echo.Context, dst interface{}) error {
	err := c.Bind(dst)
	if err != nil {
		return err
	}
	_, err = govalidator.ValidateStruct(dst)
	if err != nil {
		return err
	}
	return nil
}
