package connector

import (

	// external
	"github.com/labstack/echo"

	// product service

	"yichips/config"
	customMiddleware "yichips/middleware"
	product "yichips/services/product"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------
var mainService = product.Main{}
var productBaseUrl = config.App["product_base_url"]
var cmsApiVersion = "/v1"

/**
 * Connector
 */
func Connect(e *echo.Echo, conType string) product.Main {

	switch conType {
	case "rest":
		return InitRoutes(e)

	case "direct":
		return product.Main{}
	}

	return product.Main{}
}

func InitRoutes(e *echo.Echo) product.Main {
	// cms
	products := e.Group(productBaseUrl + cmsApiVersion + "/products")

	products.Use(customMiddleware.CMSAuthorization)

	products.GET("/list/", mainService.List)       // get product list
	products.GET("/:code/", mainService.GetDetail) // get product detail
	products.POST("/", mainService.Create)         // create new product
	// products.PUT("/:code/", Update)     // update existing product by code
	// products.DELETE("/:code/", Destroy) // delete product by code
	return product.Main{}
}
