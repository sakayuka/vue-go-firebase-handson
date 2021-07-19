package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
	}))
	e.GET("/public", public)
	e.GET("/private", private, Auth)
	e.Logger.Fatal(e.Start(":8000"))
}

func public(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello Public World!")
}

func private(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello Private World!")
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_ACCOUNT_KEY_JSON"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Println(err.Error())
			return APIResponseError(c, http.StatusInternalServerError, "invalid credential file", errors.New("invalid credential file"))
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return APIResponseError(c, http.StatusInternalServerError, err.Error(), err)
		}

		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Println(err.Error())
			return APIResponseError(c, http.StatusUnauthorized, err.Error(), err)
		}

		c.Set("token", token)
		return next(c)
	}
}

type APIErrorResponse struct {
	Code   string   `json:"code"`
	Errors []string `json:"errors"`
}

func APIResponseError(c echo.Context, httpStatus int, message string, err error) error {
	c.JSONPretty(httpStatus, APIErrorResponse{Code: fmt.Sprintf("%d-000", httpStatus), Errors: []string{message}}, "  ")
	return err
}
