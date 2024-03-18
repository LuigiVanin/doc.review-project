package helpers

import (
	e "doc-review/src/exceptions/errors"

	"github.com/go-playground/validator/v10"
)

func Validation[T interface{}](data *T) error {
	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		errorsIter := err.(validator.ValidationErrors)
		errors := make([]e.FieldError, 0)

		for _, err := range errorsIter {
			fieldError := e.FieldError{
				Field: err.Field(),
				Tag:   err.Tag(),
			}
			errors = append(errors, fieldError)
		}

		if len(errors) == 0 {
			return err
		}

		return &e.FieldErrors{
			Fields: errors,
		}
	}
	return nil
}
