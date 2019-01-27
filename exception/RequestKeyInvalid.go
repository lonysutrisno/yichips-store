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

type RequestKeyInvalid struct {
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
 * @return 	RequestKeyInvalid
 */
func NewRequestKeyInvalid() RequestKeyInvalid {
	return RequestKeyInvalid{
		Code:    "901",
		Message: "Invalid api key",
	}
}

// ------------------------------------------------------

/**
 * Get Error Code
 *
 * @return 	strinf
 */
func (ex RequestKeyInvalid) GetCode() string {
	return ex.Code
}

// ------------------------------------------------------

/**
 * Get Error Message
 *
 * @return string
 */
func (ex RequestKeyInvalid) GetMessage() string {
	return ex.Message
}
