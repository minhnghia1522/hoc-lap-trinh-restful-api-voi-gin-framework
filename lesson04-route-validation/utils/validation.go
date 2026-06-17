package utils

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ValidationRequired(fieldName string, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errorMap := make(map[string]string)

		for _, e := range validationError {
			log.Printf("%+v", e.Tag())
			log.Printf("%+v", e.Field())

			switch e.Tag() {
			case "required":
				errorMap[e.Field()] = e.Field() + " is required"
			case "gt":
				errorMap[e.Field()] = fmt.Sprintf("%s must be greater than %s", e.Field(), e.Param())
			case "lt":
				errorMap[e.Field()] = fmt.Sprintf("%s must be less than %s", e.Field(), e.Param())
			case "gte":
				errorMap[e.Field()] = fmt.Sprintf("%s must be greater than or equal to %s", e.Field(), e.Param())
			case "lte":
				errorMap[e.Field()] = fmt.Sprintf("%s must be less than or equal to %s", e.Field(), e.Param())
			case "uuid":
				errorMap[e.Field()] = e.Field() + " must be a valid UUID"
			case "slug":
				errorMap[e.Field()] = e.Field() + " must contain only lowercase letter, numbers, hyphens and dots."
			case "min":
				errorMap[e.Field()] = fmt.Sprintf("%s must be greater than %s characters.", e.Field(), e.Param())
			case "max":
				errorMap[e.Field()] = fmt.Sprintf("%s must be less than %s characters.", e.Field(), e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				errorMap[e.Field()] = fmt.Sprintf("%s must be one of %s.", e.Field(), allowedValues)
			case "search":
				errorMap[e.Field()] = e.Field() + " must contain only lowercase letter, numbers, hyphens and dots."

			}
		}
		return gin.H{"error": errorMap}
	}

	return gin.H{
		"error": "Invalid request: " + err.Error(),
	}
}

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	})

	return nil
}
