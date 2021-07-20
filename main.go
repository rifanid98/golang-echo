package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(port)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "Hello World")
		return nil
	})
	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	if err := e.Start(fmt.Sprintf(":%s", port)); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
