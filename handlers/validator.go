package handlers

import "gopkg.in/go-playground/validator.v9"

var (
	v = validator.New()
)

type requestValidator struct {
	validator *validator.Validate
}

func (u *requestValidator) Validate(i interface{}) error {
	return u.validator.Struct(i)
}
