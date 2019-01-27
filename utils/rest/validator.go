package rest

import (
	// external
	"gopkg.in/go-playground/validator.v9"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Validator struct {
	Validator *validator.Validate
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Perform Validation
 */
func (cv *Validator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
