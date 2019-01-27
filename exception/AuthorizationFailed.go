package exception

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Request Validation Failed Exception
 |	-----------------------------------
 |
 |	This exception is thrown when a request form or bas edata invalid
 |	covering mismatch Api Key, expired Timestamp, and mismatch
 |	Signature
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type AuthorizationFailed struct {
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
 * @return 	AuthorizationFailed
 */
func NewAuthorizationFailed() AuthorizationFailed {
	return AuthorizationFailed{
		Code:    "909",
		Message: "Unauthorized",
	}
}

// ------------------------------------------------------

/**
 * Get Error Code
 *
 * @return 	strinf
 */
func (ex AuthorizationFailed) GetCode() string {
	return ex.Code
}

// ------------------------------------------------------

/**
 * Get Error Message
 *
 * @return string
 */
func (ex AuthorizationFailed) GetMessage() string {
	return ex.Message
}
