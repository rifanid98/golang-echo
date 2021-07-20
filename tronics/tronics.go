package tronics

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	// method 1 echo middleware
	// will be executied before the middleware
	e.Pre(middleware.RemoveTrailingSlash())
	// method 2 echo middleware
	// will limit the size of the body that can be accepted
	e.POST("/products", createProduct, middleware.BodyLimit("1K"))
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
