package validation

import (
	"fmt"
	"regexp"
	"strings"
	"user-management-api/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func InitValidator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}
	RegisterCustomValidation(v)
	return nil
}

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errorMap := make(map[string]string)

		for _, e := range validationError {
			root := strings.Split(e.Namespace(), ".")[0]

			rawPath := strings.TrimPrefix(e.Namespace(), root+".")

			parts := strings.Split(rawPath, ".")
			for i, part := range parts {
				if strings.Contains(part, "[") {
					idx := strings.Index(part, "[")
					base := utils.CamelToSnake(part[:idx])
					index := part[idx:]
					parts[i] = base + index
				} else {
					parts[i] = utils.CamelToSnake(part)
				}
			}
			fieldPath := strings.Join(parts, ".")
			param := e.Param()
			switch e.Tag() {
			case "required":
				errorMap[fieldPath] = fmt.Sprintf("%s is required", fieldPath)
			case "gt":
				errorMap[fieldPath] = fmt.Sprintf("%s must be greater than %s", fieldPath, param)
			case "lt":
				errorMap[fieldPath] = fmt.Sprintf("%s must be less than %s", fieldPath, param)
			case "gte":
				errorMap[fieldPath] = fmt.Sprintf("%s must be greater than or equal to %s", fieldPath, param)
			case "lte":
				errorMap[fieldPath] = fmt.Sprintf("%s must be less than or equal to %s", fieldPath, param)
			case "uuid":
				errorMap[fieldPath] = fmt.Sprintf("%s must be a valid UUID", fieldPath)
			case "slug":
				errorMap[fieldPath] = fmt.Sprintf("%s must contain only lowercase letter, numbers, hyphens and dots.", fieldPath)
			case "min":
				errorMap[fieldPath] = fmt.Sprintf("%s must be greater than %s characters.", fieldPath, param)
			case "max":
				errorMap[fieldPath] = fmt.Sprintf("%s must be less than %s characters.", fieldPath, param)
			case "min_int":
				errorMap[fieldPath] = fmt.Sprintf("%s must have a greater value than %s", fieldPath, param)
			case "max_int":
				errorMap[fieldPath] = fmt.Sprintf("%s must have a value less than %s", fieldPath, param)
			case "oneof":
				allowedValues := strings.Join(strings.Split(param, " "), ", ")
				errorMap[fieldPath] = fmt.Sprintf("%s must be one of %s.", fieldPath, allowedValues)
			case "search":
				errorMap[fieldPath] = fmt.Sprintf("%s must contain only lowercase letter, numbers, hyphens and dots.", fieldPath)
			case "email":
				errorMap[fieldPath] = fmt.Sprintf("%s must be in the correct email format", fieldPath)
			case "datetime":
				errorMap[fieldPath] = fmt.Sprintf("%s must follow the YYYY-MM-DD format exactly", fieldPath)
			case "file_ext":
				allowedValues := strings.Join(strings.Split(param, " "), ",")
				errorMap[fieldPath] = fmt.Sprintf("%s only allow files with extensions: %s", fieldPath, allowedValues)
			case "email_advanced":
				errorMap[fieldPath] = fmt.Sprintf("%s is in the banned list", fieldPath)
			case "password_strong":
				errorMap[fieldPath] = fmt.Sprintf("%s must be at least 8 characters including (lowercase letters, uppercase letters, numbers, and special characters)", fieldPath)
			}
		}
		return gin.H{"error": errorMap}
	}

	return gin.H{
		"error": "Invalid request: " + err.Error(),
	}
}
