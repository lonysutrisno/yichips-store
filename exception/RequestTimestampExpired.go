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

type RequestTimestampExpired struct {
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
 * @return 	RequestTimestampExpired
 */
func NewRequestTimestampExpired() RequestTimestampExpired {
	return RequestTimestampExpired{
		Code:    "902",
		Message: "Request Expired",
	}
}

// ------------------------------------------------------

/**
 * Get Error Code
 *
 * @return 	strinf
 */
func (ex RequestTimestampExpired) GetCode() string {
	return ex.Code
}

// ------------------------------------------------------

/**
 * Get Error Message
 *
 * @return string
 */
func (ex RequestTimestampExpired) GetMessage() string {
	return ex.Message
}
