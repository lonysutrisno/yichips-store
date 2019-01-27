package output

import (
	// yichips
	"yichips/utils/contract"

	// authcms service
	"yichips/services/user/model"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type UserLogin struct {
	ApiToken           string `json:"api_token"`
	ApiTokenExpiration int    `json:"expiration"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * User Login Output Model
 *
 * @param 	*model.User 	user
 *
 * @return 	contract.ModelResponse
 */
func NewUserLogin(user *model.User) contract.ModelResponse {
	return UserLogin{
		ApiToken:           user.ApiToken,
		ApiTokenExpiration: int(user.ApiTokenExpiration.Unix()),
	}
}
