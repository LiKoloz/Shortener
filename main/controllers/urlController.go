package controllers

import (
	"Shortener/models"
	"Shortener/services"

	"github.com/labstack/echo"
)

func UrlController(e *echo.Echo) {

	e.POST("/shorten", func(c echo.Context) error {
		url := new(models.Url)
		if err := c.Bind(url); err != nil {
			return c.String(400, "Uncorrect url")
		}
		err := services.Add(*url)
		if err != nil {
			return c.String(500, "Smth wrong")
		}
		return c.NoContent(200)
	})

	e.GET("/s/:short_url", func(c echo.Context) error {
		shortUrl := c.Param("short_url")
		result, err := services.GetLongByShort(shortUrl)
		if err != nil {
			return c.String(500, "Smth wrong")
		}
		return c.String(200, result)
	})

	e.GET("/analytics/:short_url", func(c echo.Context) error {
		shortUrl := c.Param("short_url")
		result, err := services.GetByShort(shortUrl)
		if err != nil {
			return c.String(500, "Smth wrong")
		}
		return c.JSON(200, result)
	})

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.File("static/index.html")
	})
}
