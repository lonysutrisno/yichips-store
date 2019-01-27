package model

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type ApiErrorResponse struct {
	Status    string `json:"status"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	Signature string `json:"signature"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTION
// ------------------------------------------------------

/**
 * Add timestamp and sign response
 *
 * @param 	*ApiResponse 	resp
 */
func (resp *ApiErrorResponse) AddTimestampAndSign() {
	resp.Timestamp = AddTimestamp()
	resp.Signature = Sign(resp)
}
