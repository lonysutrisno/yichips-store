package rest

import (
	// native
	"net/http"

	// external
	"github.com/labstack/echo"

	// yichips
	"yichips/exception"
	"yichips/utils/contract"
	"yichips/utils/rest/model"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Construct API Error Response
 */
func ConstructErrorResponse(e echo.Context, ex exception.Exception) error {
	response := model.ApiErrorResponse{
		Status:  "ERROR",
		Code:    ex.GetCode(),
		Message: ex.GetMessage(),
	}

	response.AddTimestampAndSign()

	return e.JSON(http.StatusBadRequest, response)
}

// ------------------------------------------------------

/**
 * Construct API Success Response
 */
func ConstructSuccessResponse(e echo.Context, data map[string]contract.Model) error {
	response := model.ApiSuccessResponse{
		Status: "SUCCESS",
		Code:   "000",
		Data:   data,
	}

	response.AddTimestampAndSign()

	return e.JSON(http.StatusOK, response)
}
