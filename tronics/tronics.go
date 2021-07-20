package tronics

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

var (
	e = echo.New()
	v = validator.New()
)

// Start starts the application
func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.POST("/products", createProduct)
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
