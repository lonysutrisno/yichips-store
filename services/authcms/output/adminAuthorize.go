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

type AdminAuthorize struct {
	ID     uint
	Code   string
	Name   string
	Email  string
	Status string
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Admin Detail Output Model
 *
 * @param 	*model.Admin admin
 *
 * @return 	contract.ModelResponse
 */
func NewAdminAuthorize(admin *model.Admin) contract.ModelResponse {
	return AdminAuthorize{
		ID:     admin.ID,
		Code:   admin.Code,
		Name:   admin.Name,
		Email:  admin.Email,
		Status: admin.Status,
	}
}
