package cms

import (
	// external
	"github.com/labstack/echo"

	// yichips
	"yichips/exception"
	"yichips/utils/contract"
	"yichips/utils/rest"

	// authcms service
	"yichips/services/authcms"
	"yichips/services/authcms/connector/rest/cms/request"

	// other services
	admin "yichips/services/admin/connector"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Controller struct {
}

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------

var Version = "v1"
var mainService = authcms.Main{}
var adminService = admin.Connect(nil, "direct")

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------
//

/**
 * Handle Admin Login
 */
func (ctr *Controller) Login(e echo.Context) (err error) {
	// bind and validate request
	req := new(request.Login)
	e.Bind(req)
	if err = e.Validate(req); err != nil {
		return rest.ConstructErrorResponse(e, exception.NewInputValidationFailed(err.Error()))
	}

	// attempt to login admin
	authorization, adminCode, exc := mainService.Login(req.Email, req.Password)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// get admin details
	adminDetail, exc := adminService.GetDetail(adminCode)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	data := map[string]contract.Model{
		"authorization": authorization,
		"admin":         adminDetail,
	}

	return rest.ConstructSuccessResponse(e, data)
}
