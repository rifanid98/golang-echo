package tronics

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (pv *ProductValidator) Validate(i interface{}) error {
	return pv.validator.Struct(i)
}

var (
	products = []map[int]string{
		{1: "Laptop"},
		{2: "Phone"},
		{3: "Mouse"},
	}
)

func getProducts(c echo.Context) error {
	e.Logger.Print("Inside of getProducts")
	query := strings.TrimSpace(c.QueryParam("name"))
	if query == "" {
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusNotFound, "name query param was disabled")
	}
	return nil
}

func getProduct(c echo.Context) error {
	var product map[int]string

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for _, p := range products {
		for k := range p {
			if productId == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	type body struct {
		Name string `json:"name" validate:"required,min=4"`
	}

	var reqBody body
	e.Validator = &ProductValidator{validator: (*validator.Validate)(v)}
	if err := c.Bind(&reqBody); err != nil {
		e.Logger.Fatal(err)
		return err
	}
	if err := c.Validate(&reqBody); err != nil {
		return err
	}

	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}

	products = append(products, product)
	return c.JSON(http.StatusCreated, product)
}

func updateProduct(c echo.Context) error {
	var product map[int]string

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for _, p := range products {
		for k := range p {
			if productId == k {
				product = p
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	type body struct {
		Name string `json:"name" validate:"required,min=4"`
	}

	var reqBody body
	e.Validator = &ProductValidator{validator: (*validator.Validate)(v)}
	if err := c.Bind(&reqBody); err != nil {
		e.Logger.Fatal(err)
		return err
	}
	if err := c.Validate(&reqBody); err != nil {
		return err
	}

	product[productId] = reqBody.Name

	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	var product map[int]string

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var index int
	for i, p := range products {
		for k := range p {
			if productId == k {
				product = p
				index = i
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	products = append(products[:index], products[index+1:]...)

	return c.JSON(http.StatusOK, product)
}
