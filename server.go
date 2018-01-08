package main

import (
	"net/http"

	"github.com/iamgoangle/go-crud/user"
	"github.com/labstack/echo"
)

// Declare server configuration
const port string = ":3000"

// Router default router of my app
func Router() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/createUser", user.CreateUser)
	e.GET("/users/:id", user.GetUser)
	e.GET("users", user.GetUsers)
	e.GET("/show", user.QueryUser)

	// e.PUT("/users/:id", updateUser)

	// e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(port))
}

func main() {
	Router()
}
