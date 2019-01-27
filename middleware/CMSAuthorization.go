package middleware

import (
	// native
	"os"
	"strings"

	// external
	"github.com/labstack/echo"

	// yichips
	"yichips/exception"
	"yichips/utils/rest"

	// authcms service
	"yichips/services/authcms"
)

/**
 * Authorize API Request Actor
 *
 * @param 	echo.HandleFunc 	next
 *
 * @return 	echo.HandleFunc
 */
func CMSAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		if os.Getenv("SKIP_AUTHORIZATION_CHECK") != "true" {
			// test Authorization header existence
			if e.Request().Header["Authorization"] == nil {
				return rest.ConstructErrorResponse(e, exception.NewAuthorizationFailed())
			}

			// test whether Authorization implement correct format
			apiTokenRaw := e.Request().Header["Authorization"][0]
			apiTokenSplit := strings.Split(apiTokenRaw, " ")
			if len(apiTokenSplit) != 2 || apiTokenSplit[0] != "Bearer" {
				return rest.ConstructErrorResponse(e, exception.NewAuthorizationFailed())
			}

			// attempt to authorize
			service := authcms.Main{}
			actor, exc := service.Authorize(apiTokenSplit[1])
			if exc != nil {
				return rest.ConstructErrorResponse(e, exc)
			}

			// set request context _actor (authorized user)
			e.Set("_actor", actor)
		}

		if err := next(e); err != nil {
			e.Error(err)
		}

		return nil
	}
}
