package main

import (
	"net/http"
	"strategy-pattern/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/checkout", handler.CalculatePrice)
	e.Logger.Fatal(e.Start(":1323"))
}
