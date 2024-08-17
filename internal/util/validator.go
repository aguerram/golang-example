package util

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mostafa/learn_go/internal/type/global"
	"reflect"
	"strings"
)

func getJSONFieldName(structType reflect.Type, fieldName string) string {
	if field, ok := structType.FieldByName(fieldName); ok {
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			return strings.Split(jsonTag, ",")[0] // In case there are options like `json:"username,omitempty"`
		}
	}
	return fieldName
}

func ValidateRequest(r any, validator *global.XValidator) *fiber.Error {
	if errs := validator.Validate(r); len(errs) > 0 && errs[0].Error {
		errorMessages := make([]string, 0)

		val := reflect.ValueOf(r)
		structType := val.Type()

		for _, err := range errs {
			errorMessages = append(errorMessages, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				getJSONFieldName(structType, err.FailedField),
				err.Value,
				err.Tag,
			))
		}

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errorMessages, " and "),
		}
	}
	return nil
}
