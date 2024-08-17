package global

import (
	"github.com/go-playground/validator/v10"
	"mostafa/learn_go/internal/type/response"
)

type XValidator struct {
	validator *validator.Validate
}

func NewValidator() *XValidator {
	return &XValidator{
		validator: validator.New(),
	}
}

func (v XValidator) Validate(data interface{}) []response.ErrorResponse {
	var validationErrors []response.ErrorResponse

	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem response.ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
