package utils

import (
	"fmt"
	"log"
	"regexp"

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
			case "gt":
				errorMap[e.Field()] = e.Field() + " must be a positive number"
			case "uuid":
				errorMap[e.Field()] = e.Field() + " must be a valid UUID"
			case "slug":
				errorMap[e.Field()] = e.Field() + " must contain only lowercase letter, numbers, hyphens and dots."
			case "min":
				errorMap[e.Field()] = fmt.Sprintf("%s must be greater than %s characters.", e.Field(), e.Param())
			case "max":
				errorMap[e.Field()] = fmt.Sprintf("%s must be less than %s characters.", e.Field(), e.Param())
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
	return nil
}
