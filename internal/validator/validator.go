package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		var errMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errMessages = append(errMessages, err.Error())
		}
		return errors.New(formatValidationErrors(err))
	}

	return nil
}

func formatValidationErrors(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var errMessages []string
		for _, e := range errs {
			errMessages = append(errMessages, fmt.Sprintf("Field '%s' %s", e.Field(), e.Tag()))
		}

		return strings.Join(errMessages, ", ")
	}

	return err.Error()
}
