package helper

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsValidationError(err error) bool {
	var validationErrors validator.ValidationErrors
	return errors.As(err, &validationErrors)
}
