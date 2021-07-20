package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "Hello World")
		return nil
	})
	products := []map[int]string{
		{1: "Laptop"},
		{2: "Phone"},
		{3: "Mouse"},
	}
	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string
		for _, p := range products {
			for k := range p {
				productId, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if productId == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}
		return c.JSON(http.StatusOK, product)
	})
	e.Logger.Print("Listening on port 8080")
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
