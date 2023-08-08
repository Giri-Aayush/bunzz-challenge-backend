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

	port := viper.GetString("PORT")
	if port == "" {
		port = "3000"
	}

	if err := e.Start("0.0.0.0:" + port); err != nil {
		fmt.Printf("Error while starting server %s", err)
	}

}
