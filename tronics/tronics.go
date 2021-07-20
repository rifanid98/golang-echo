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

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		e.Logger.Print("inside serverHeader middleware")
		return next(c)
	}
}

// Start starts the application
func Start() {
	// method 1 to use middleware
	e.Use(serverMessage)
	e.POST("/products", createProduct)
	// method 2 to use middleware
	e.GET("/products", getProducts, serverMessage)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
