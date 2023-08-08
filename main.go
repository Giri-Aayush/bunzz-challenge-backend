package main

import (
	"fmt"

	"github.com/Giri-Aayush/bunzz-challenge-backend/handlers"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error while reading config file %s", err)
	}
}

func main() {
	e := echo.New()
	h := &handlers.FizzBuzzHandler{}

	e.POST("/fizzbuzz", h.GetMessage)

	if err := e.Start("localhost:3000"); err != nil {
		fmt.Printf("Error while starting server %s", err)
	}

}
