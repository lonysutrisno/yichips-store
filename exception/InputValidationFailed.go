package exception

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Request Validation Failed Exception
 |	-----------------------------------
 |
 |	This exception is thrown when a request data received fails to
 |	validate the value based on the given rule
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type InputValidationFailed struct {
	Code    string
	Message string
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Constructor
 *
 * @return 	AccountInactive
 */
func NewInputValidationFailed(message string) InputValidationFailed {
	return InputValidationFailed{
		Code:    "900",
		Message: message,
	}
}

// ------------------------------------------------------

/**
 * Get Error Code
 *
 * @return 	strinf
 */
func (ex InputValidationFailed) GetCode() string {
	return ex.Code
}

// ------------------------------------------------------

/**
 * Get Error Message
 *
 * @return string
 */
func (ex InputValidationFailed) GetMessage() string {
	return ex.Message
}
