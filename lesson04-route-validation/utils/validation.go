package utils

import "fmt"

func ValidationRequired(fieldName string, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}
