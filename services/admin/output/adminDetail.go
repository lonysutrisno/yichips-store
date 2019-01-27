package output

import (
	// yichips
	"yichips/utils/collection"
	"yichips/utils/contract"

	// admin service
	"yichips/services/admin/model"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type AdminDetail struct {
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
 * Admin Detail Output Model
 *
 * @param 	*model.Admin 	admin
 *
 * @return 	contract.ModelResponse
 */
func NewAdminDetail(admin *model.Admin) contract.ModelResponse {
	return AdminDetail{
		Code:   admin.Code,
		Name:   admin.Name,
		Email:  admin.Email,
		Status: admin.Status,
	}
}

// ------------------------------------------------------

/**
 * Admin List Output Models
 *
 * @param 	[]model.Admin 	admin
 *
 * @return 	collection.ModelResponse
 */
func NewAdminDetailList(adminList []model.Admin) collection.ModelResponse {
	adminDetails := collection.NewModelResponse()
	for _, admin := range adminList {
		adminDetails.Add(NewAdminDetail(&admin))
	}

	return adminDetails
}
