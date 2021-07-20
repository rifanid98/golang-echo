package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	e *echo.Echo = echo.New()
)

func main() {
	mainRoute()
	productsRoute()
	startEcho()
}

func startEcho() {
	port := "8080"
	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	if err := e.Start(fmt.Sprintf(":%s", port)); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}

func mainRoute() *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "Hello World")
		return nil
	})
	return e
}

func productsRoute() *echo.Echo {
	products := []map[int]string{
		{1: "Laptop"},
		{2: "Phone"},
		{3: "Mouse"},
	}

	e.GET("/products/", func(c echo.Context) error {
		c.JSON(http.StatusOK, products)
		return nil
	})

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

	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"name"`
		}

		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			e.Logger.Fatal(err)
			return err
		}

		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}

		products = append(products, product)
		return c.JSON(http.StatusCreated, product)
	})
	return e
}
