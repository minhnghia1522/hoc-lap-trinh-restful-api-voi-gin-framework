package utils

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
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
			fieldPath := e.Field()
			param := e.Param()
			switch e.Tag() {
			case "required":
				errorMap[fieldPath] = fieldPath + " is required"
			case "gt":
				errorMap[fieldPath] = fmt.Sprintf("%s must be greater than %s", fieldPath, param)
			case "lt":
				errorMap[fieldPath] = fmt.Sprintf("%s must be less than %s", fieldPath, param)
			case "gte":
				errorMap[fieldPath] = fmt.Sprintf("%s must be greater than or equal to %s", fieldPath, param)
			case "lte":
				errorMap[fieldPath] = fmt.Sprintf("%s must be less than or equal to %s", fieldPath, param)
			case "uuid":
				errorMap[fieldPath] = fieldPath + " must be a valid UUID"
			case "slug":
				errorMap[fieldPath] = fieldPath + " must contain only lowercase letter, numbers, hyphens and dots."
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
				errorMap[fieldPath] = fieldPath + " must contain only lowercase letter, numbers, hyphens and dots."
			case "email":
				errorMap[fieldPath] = fmt.Sprintf("%s must be in the correct email format", fieldPath)
			case "datetime":
				errorMap[fieldPath] = fmt.Sprintf("%s must follow the YYYY-MM-DD format exactly", fieldPath)
			case "file_ext":
				allowedValues := strings.Join(strings.Split(param, " "), ",")
				errorMap[fieldPath] = fmt.Sprintf("%s only allow files with extensions: %s", fieldPath, allowedValues)
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

	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		minVal, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= minVal
	})

	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		maxVal, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() <= maxVal
	})

	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()

		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}

		allowedExt := strings.Fields(allowedStr)
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowedExt {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}

		return false
	})

	return nil
}
