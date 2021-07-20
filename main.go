package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	e *echo.Echo = echo.New()
	v            = validator.New()
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

	e.GET("/products", func(c echo.Context) error {
		query := strings.TrimSpace(c.QueryParam("name"))
		if query == "" {
			c.JSON(http.StatusOK, products)
		} else {
			c.JSON(http.StatusNotFound, "name query param was disabled")
		}
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
			Name string `json:"name" validate:"required,min=4"`
			// Vendor          string `json:"vendor" validate:"min=5,max=10"`
			// Email           string `json:"email" validate:"required_with=Vendor,email"`
			// Website         string `json:"website" validagte:"url"`
			// Country         string `json:"country" validate:"ip"`
			// DefaultDeviceIp string `json:"default_device_ip" validate:"ip"`
		}

		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			e.Logger.Fatal(err)
			return err
		}
		if err := v.Struct(&reqBody); err != nil {
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
