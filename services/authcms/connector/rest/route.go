package rest

import (
	// external
	"github.com/labstack/echo"

	// fitce
	"yichips/config"

	// authcms service
	"yichips/services/authcms"
	"yichips/services/authcms/connector/rest/cms"
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
func InitRoutes(e *echo.Echo) authcms.Main {
	// cms
	e.POST(adminBaseUrl+cmsApiVersion+"/login/", cmsCtr.Login) // login admin

	return authcms.Main{}
}
