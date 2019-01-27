package output

import (
	// yichips
	"yichips/utils/contract"

	// authcms service
	"yichips/services/authcms/model"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type AdminLogin struct {
	ApiToken           string `json:"api_token"`
	ApiTokenExpiration int    `json:"expiration"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Admin Login Output Model
 *
 * @param 	*model.Admin 	admin
 *
 * @return 	contract.ModelResponse
 */
func NewAdminLogin(admin *model.Admin) contract.ModelResponse {
	return AdminLogin{
		ApiToken:           admin.ApiToken,
		ApiTokenExpiration: int(admin.ApiTokenExpiration.Unix()),
	}
}
