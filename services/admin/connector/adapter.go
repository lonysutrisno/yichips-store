package connector

import (
	// external
	"github.com/labstack/echo"

	// admin service
	"yichips/services/admin"
	"yichips/services/admin/connector/rest"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Connector
 */
func Connect(e *echo.Echo, conType string) admin.Main {
	switch conType {
	case "rest":
		return rest.InitRoutes(e)

	case "direct":
		return admin.Main{}
	}

	return admin.Main{}
}
