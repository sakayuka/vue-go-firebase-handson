package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
	}))
	e.GET("/public", public)
	e.GET("/private", private)
	e.Logger.Fatal(e.Start(":8000"))
}

func public(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Public World!")
}

func private(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Private World!")
}
