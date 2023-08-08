package main

import (
	"fmt"
	"os"

	"github.com/Giri-Aayush/bunzz-challenge-backend/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	h := &handlers.FizzBuzzHandler{}

	// Add CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.POST("/fizzbuzz", h.GetMessage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := e.Start("0.0.0.0:" + port); err != nil {
		fmt.Printf("Error while starting server %s", err)
	}

}
