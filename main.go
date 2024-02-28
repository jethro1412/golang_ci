package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jethro1412/golang_ci/api"

	_ "github.com/jethro1412/golang_ci/docs" // This line is necessary for go-swagger to find your docs!
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if username == "techincal" && password == "test" {
				return true, nil
			}
			return false, nil
		},
	}))

	// Route => handler
	e.POST("/api", api.ApiHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}