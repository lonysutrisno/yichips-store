package middleware

import (
	// native
	"os"
	"strconv"
	"time"

	// yichips
	"yichips/exception"
	"yichips/utils/rest"

	// external
	"github.com/labstack/echo"
)

type BaseRequestData struct {
	Timestamp int    `json:"timestamp"`
	Signature string `json:"signature"`
}

/**
 * Validate Request Base
 *
 * @param 	echo.HandleFunc 	next
 *
 * @return 	echo.HandleFunc
 */
func RequestValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		if os.Getenv("SKIP_API_KEY_CHECK") != "true" {
			// test X-Api-Key header
			if e.Request().Header["X-Api-Key"] == nil || e.Request().Header["X-Api-Key"][0] != os.Getenv("API_KEY") {
				return rest.ConstructErrorResponse(e, exception.NewRequestKeyInvalid())
			}
		}

		if os.Getenv("SKIP_TIMESTAMP_CHECK") != "true" {
			// test request expiration
			maxMinute, _ := strconv.Atoi(os.Getenv("REQUEST_EXPIRATION_DURATION"))
			allowedTimestamp := int(time.Now().Add(time.Minute * -time.Duration(maxMinute)).Unix())
			// if e.Request().Method == "GET" {
			requestTimestamp, err := strconv.Atoi(e.QueryParam("timestamp"))
			if err != nil || requestTimestamp < allowedTimestamp {
				return rest.ConstructErrorResponse(e, exception.NewRequestTimestampExpired())
			}
			// } else {
			// 	req := new(BaseRequestData)
			// 	e.Bind(req)
			// 	requestTimestamp := req.Timestamp
			// 	if requestTimestamp < allowedTimestamp {
			// 		return rest.ConstructErrorResponse(e, exception.NewRequestValidationFailed("Request Expired"))
			// 	}
			// }
		}

		if err := next(e); err != nil {
			e.Error(err)
		}

		return nil
	}
}
