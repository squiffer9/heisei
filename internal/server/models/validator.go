package models

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

// init initializes the validator.
func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct.
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
