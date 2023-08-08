package main

import (
	"fmt"
	"os"

	"github.com/Giri-Aayush/bunzz-challenge-backend/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	h := &handlers.FizzBuzzHandler{}

	e.POST("/fizzbuzz", h.GetMessage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := e.Start("0.0.0.0:" + port); err != nil {
		fmt.Printf("Error while starting server %s", err)
	}

}
