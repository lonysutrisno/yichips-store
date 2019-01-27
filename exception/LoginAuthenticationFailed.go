package exception

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Login Authentication Failed Exception
 |	-------------------------------------
 |
 |	This exception is thrown when a user/member/admin attempts to do
 |	login/authentication but failed while trying to prove to do so
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type LoginAuthenticationFailed struct {
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
func NewLoginAuthenticationFailed() LoginAuthenticationFailed {
	return LoginAuthenticationFailed{
		Code:    "201",
		Message: "Email or Password is incorrect",
	}
}

// ------------------------------------------------------

/**
 * Get Error Code
 *
 * @return 	strinf
 */
func (ex LoginAuthenticationFailed) GetCode() string {
	return ex.Code
}

// ------------------------------------------------------

/**
 * Get Error Message
 *
 * @return string
 */
func (ex LoginAuthenticationFailed) GetMessage() string {
	return ex.Message
}
