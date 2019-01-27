package rest

import (
	// external

	"github.com/labstack/echo"

	// yichips
	"yichips/config"
	customMiddleware "yichips/middleware"

	// admin service
	"yichips/services/admin"
	"yichips/services/admin/connector/rest/cms"
)

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------

var cmsCtr = cms.Controller{}
var adminBaseUrl = config.App["admin_base_url"]
var cmsApiVersion = "/" + cms.Version

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Rout Sets
 */
func InitRoutes(e *echo.Echo) admin.Main {
	// cms
	admins := e.Group(adminBaseUrl + cmsApiVersion + "/admins")

	admins.Use(customMiddleware.CMSAuthorization)

	admins.GET("/", cmsCtr.Index)            // list admin
	admins.GET("/:code/", cmsCtr.Show)       // get admin by code
	admins.POST("/", cmsCtr.Create)          // create new admin
	admins.PUT("/:code/", cmsCtr.Update)     // update existing admin by code
	admins.DELETE("/:code/", cmsCtr.Destroy) // delete admin by code

	return admin.Main{}
}
