package connector

import (

	// external
	"github.com/labstack/echo"

	// user service

	"yichips/config"
	customMiddleware "yichips/middleware"
	"yichips/services/user"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------
var mainService = user.Main{}
var userBaseUrl = config.App["user_base_url"]
var cmsApiVersion = "/v1"

/**
 * Connector
 */
func Connect(e *echo.Echo, conType string) user.Main {

	switch conType {
	case "rest":
		return InitRoutes(e)

	case "direct":
		return user.Main{}
	}

	return user.Main{}
}
func InitRoutes(e *echo.Echo) user.Main {
	// cms
	e.POST(userBaseUrl+cmsApiVersion+"/users/login/", mainService.LoginByMail) // login user
	e.GET(userBaseUrl+cmsApiVersion+"/users/verification/:token/", mainService.MailVerification)
	users := e.Group(userBaseUrl + cmsApiVersion + "/users")

	users.Use(customMiddleware.CMSAuthorization)

	users.POST("/register/", mainService.Register) // register user

	users.GET("/detail/:code/", mainService.GetDetail) // get user by code
	// verification by mail
	// users.POST("/", Create)          // create new user
	// users.PUT("/:code/", Update)     // update existing user by code
	// users.DELETE("/:code/", Destroy) // delete user by code
	return user.Main{}
}
