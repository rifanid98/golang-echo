package tronics

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
)

var (
	e = echo.New()
	v = validator.New()
)

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

// Start starts the application
func Start() {
	e.POST("/products", createProduct)
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
