package main

import (
	"Shortener/controllers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	controllers.UrlController(e)

	if err := e.Start(":8081"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
