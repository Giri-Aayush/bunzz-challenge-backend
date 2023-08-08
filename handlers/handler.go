package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type FizzbuzzBody struct {
	Count int `json:"count" validate:"required"` //payload should have count
}

type FizzBuzzHandler struct {
}

func (h *FizzBuzzHandler) GetMessage(c echo.Context) error {
	var requestBody FizzbuzzBody

	c.Echo().Validator = &requestValidator{validator: v}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusInternalServerError, DefaultResponse{
			Status:  500,
			Message: "Internal server error",
		})
	}

	if err := c.Validate(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, DefaultResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	return c.JSON(http.StatusOK, DefaultResponse{
		Status:  200,
		Message: messageCondition(requestBody.Count),
	})
}

func messageCondition(cnt int) string {
	fizzMessage := viper.GetString("FIZZ_MESSAGE")
	buzzMessage := viper.GetString("BUZZ_MESSAGE")
	fizzBuzzMessage := viper.GetString("FIZZBUZZ_MESSAGE")

	switch {
	case cnt%3 == 0 && cnt%5 == 0:
		return fizzBuzzMessage
	case cnt%3 == 0:
		return fizzMessage
	case cnt%5 == 0:
		return buzzMessage
	default:
		return ""
	}
}
