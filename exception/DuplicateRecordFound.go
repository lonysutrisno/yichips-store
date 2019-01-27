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

type DuplicateRecordFound struct {
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
func NewDuplicateRecordFound(recordName string) DuplicateRecordFound {
	if recordName == "" {
		recordName = "Record"
	}

	return DuplicateRecordFound{
		Code:    "100",
		Message: recordName + " is already exist",
	}
}

// ------------------------------------------------------

/**
 * Get Error Code
 *
 * @return 	strinf
 */
func (ex DuplicateRecordFound) GetCode() string {
	return ex.Code
}

// ------------------------------------------------------

/**
 * Get Error Message
 *
 * @return string
 */
func (ex DuplicateRecordFound) GetMessage() string {
	return ex.Message
}
