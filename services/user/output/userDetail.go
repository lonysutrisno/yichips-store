package output

import (
	// yichips
	"yichips/utils/collection"
	"yichips/utils/contract"

	// user service
	"yichips/services/user/model"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type UserDetail struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * User Detail Output Model
 *
 * @param 	*model.User 	user
 *
 * @return 	contract.ModelResponse
 */
func NewUserDetail(user *model.User) contract.ModelResponse {
	return UserDetail{
		Code:   user.Code,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
}

// ------------------------------------------------------

/**
 * User List Output Models
 *
 * @param 	[]model.User 	user
 *
 * @return 	collection.ModelResponse
 */
func NewUserDetailList(userList []model.User) collection.ModelResponse {
	userDetails := collection.NewModelResponse()
	for _, user := range userList {
		userDetails.Add(NewUserDetail(&user))
	}

	return userDetails
}
