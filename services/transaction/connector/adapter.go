package connector

import (

	// external
	"github.com/labstack/echo"

	// transaction service

	"yichips/config"
	customMiddleware "yichips/middleware"
	transaction "yichips/services/transaction"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------
var mainService = transaction.Main{}
var transactionBaseUrl = config.App["transaction_base_url"]
var cmsApiVersion = "/v1"

/**
 * Connector
 */
func Connect(e *echo.Echo, conType string) transaction.Main {

	switch conType {
	case "rest":
		return InitRoutes(e)

	case "direct":
		return transaction.Main{}
	}

	return transaction.Main{}
}

func InitRoutes(e *echo.Echo) transaction.Main {
	// cms
	transactions := e.Group(transactionBaseUrl + cmsApiVersion + "/transactions")

	transactions.Use(customMiddleware.CMSAuthorization)

	// transactions.GET("/list/", mainService.List)       // get transaction list
	// transactions.GET("/:code/", mainService.GetDetail) // get transaction detail
	transactions.POST("/", mainService.Create) // create new transaction
	// transactions.PUT("/:code/", Update)     // update existing transaction by code
	// transactions.DELETE("/:code/", Destroy) // delete transaction by code
	return transaction.Main{}
}
