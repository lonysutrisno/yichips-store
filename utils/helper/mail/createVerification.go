package mail

import (
	"encoding/base64"

	"github.com/labstack/echo"
)

/**
 * Generate Random String with Defined Length
 *
 * @params 	int 	length
 *
 * @return 	string
 */
func CreateLink(code string, c echo.Context) string {
	// initialize set params
	r := c.Request()
	baseURL := c.Scheme() + "://" + r.Host
	str := base64.StdEncoding.EncodeToString([]byte(code))
	return baseURL + "/v1/users/verification/" + str
}
