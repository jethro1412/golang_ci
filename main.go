package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jethro1412/golang_ci/api"
	"github.com/swaggo/echo-swagger"
	_ "github.com/jethro1412/golang_ci/docs" // This line is necessary for go-swagger to find your docs!
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the APP")
}


func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if username == "technical" && password == "test" {
				return true, nil
			}
			return false, nil
		},
	}))

	// Route => handler
	e.GET("/", hello)
	e.POST("/api", api.ApiHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}