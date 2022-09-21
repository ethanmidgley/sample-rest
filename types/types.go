// Package types holds any types that may be required in multiple packages
package types

// FieldError is a struct which will be returned when an input fails validation
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
