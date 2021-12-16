package domain

import (
	"errors"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type Member struct {
	Name            string `json:"name" validate:"alpha|containsany= ."`
	Email           string `json:"email" validate:"email"`
	RegistationDate int64  `json:"registration_date"`
}

func (member Member) Validate() (err error) {
	var errorMessage string
	v := validator.New()
	validateErr := v.Struct(member)
	if validateErr != nil {
		fieldError := validateErr.(validator.ValidationErrors)
		for _, el := range fieldError {
			switch el.Tag() {
			case "alpha|containsany= .":
				errorMessage += fmt.Sprintf(
					"field %s must contains English letters, dots, spaces", el.Field())
			case "email":
				errorMessage += fmt.Sprintf(
					"field %s must contains valid email format", el.Field())
			default:
				errorMessage += fmt.Sprintf(
					"something wrong with %s validation", el.Field())
			}
		}
		err = errors.New(errorMessage)
		return
	}
	return
}
