package connector

import (

	// external
	"github.com/labstack/echo"

	// cart service

	"yichips/config"
	customMiddleware "yichips/middleware"
	cart "yichips/services/cart"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------
var mainService = cart.Main{}
var cartBaseUrl = config.App["cart_base_url"]
var cmsApiVersion = "/v1"

/**
 * Connector
 */
func Connect(e *echo.Echo, conType string) cart.Main {

	switch conType {
	case "rest":
		return InitRoutes(e)

	case "direct":
		return cart.Main{}
	}

	return cart.Main{}
}

func InitRoutes(e *echo.Echo) cart.Main {
	// cms
	carts := e.Group(cartBaseUrl + cmsApiVersion + "/carts")

	carts.Use(customMiddleware.CMSAuthorization)

	// carts.GET("/list/", mainService.List)       // get cart list
	carts.GET("/:code/", mainService.GetCartByUserId)  // get cart by user ID
	carts.POST("/", mainService.Create)                // create new cart
	carts.PUT("/:code/", mainService.UpdateCart)       // update existing cart by code
	carts.DELETE("/:code/", mainService.RemoveProduct) // delete cart by code
	carts.DELETE("/all/:code/", mainService.EmptyCart) // delete cart by code
	return cart.Main{}
}
