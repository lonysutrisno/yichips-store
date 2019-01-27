package exception

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Record Not Found Exception
 |	--------------------------
 |
 |	This exception is thrown when a requested data cannot be found
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type RecordNotFound struct {
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
 * @param 	string 	recordName
 *
 * @return 	AccountInactive
 */
func NewRecordNotFound(recordName string) RecordNotFound {
	if recordName == "" {
		recordName = "Record"
	}

	return RecordNotFound{
		Code:    "100",
		Message: recordName + " does not exist",
	}
}

// ------------------------------------------------------

/**
 * Get Error Code
 *
 * @return 	strinf
 */
func (ex RecordNotFound) GetCode() string {
	return ex.Code
}

// ------------------------------------------------------

/**
 * Get Error Message
 *
 * @return string
 */
func (ex RecordNotFound) GetMessage() string {
	return ex.Message
}
