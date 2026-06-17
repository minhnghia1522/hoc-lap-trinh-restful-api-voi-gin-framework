package utils

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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
			}

		}
		return gin.H{"error": errorMap}
	}

	return gin.H{
		"error": "Invalid request: " + err.Error(),
	}
}
