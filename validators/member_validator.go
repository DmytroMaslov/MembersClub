package validators

import (
	"errors"
	"fmt"

	"github.com/DmytroMaslov/memberclub/domain"
	"gopkg.in/go-playground/validator.v9"
)

func Validate(member domain.Member) (err error) {
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
			case "lte":
				errorMessage += fmt.Sprintf(
					"field %s must contains creative data less than now", el.Field())
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
