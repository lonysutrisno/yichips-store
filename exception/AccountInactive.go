package exception

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Account Inactive Exception
 |	--------------------------
 |
 |	This exception is thrown when a requested user/member/admin is in
 |	inactive mode/status when it expected to be in active mode/status
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type AccountInactive struct {
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
func NewAccountInactive() AccountInactive {
	return AccountInactive{
		Code:    "202",
		Message: "Account is inactive",
	}
}

// ------------------------------------------------------

/**
 * Get Error Code
 *
 * @return 	strinf
 */
func (ex AccountInactive) GetCode() string {
	return ex.Code
}

// ------------------------------------------------------

/**
 * Get Error Message
 *
 * @return string
 */
func (ex AccountInactive) GetMessage() string {
	return ex.Message
}
