package validator

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatValidationErrors(err error) map[string]string {
	errMap := make(map[string]string)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {

			// get JSON field name
			field := strings.ToLower(e.Field())

			switch e.Tag() {
			case "required":
				errMap[field] = field + " is required"
			default:
				errMap[field] = field + " is invalid"
			}
		}
	}

	return errMap
}
