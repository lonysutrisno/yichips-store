package model

import (
	// yichips
	"yichips/utils/contract"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type ApiSuccessResponse struct {
	Status    string                    `json:"status"`
	Code      string                    `json:"code"`
	Data      map[string]contract.Model `json:"data"`
	Timestamp int                       `json:"timestamp"`
	Signature string                    `json:"signature"`
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
func (resp *ApiSuccessResponse) AddTimestampAndSign() {
	resp.Timestamp = AddTimestamp()
	resp.Signature = Sign(resp)
}
