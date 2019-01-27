package connector

import (
	// external
	"github.com/labstack/echo"

	// authcms service
	"yichips/services/authcms"
	"yichips/services/authcms/connector/rest"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Connector
 */
func Connect(e *echo.Echo, conType string) authcms.Main {
	switch conType {
	case "rest":
		return rest.InitRoutes(e)

	case "direct":
		return authcms.Main{}
	}

	return authcms.Main{}
}
