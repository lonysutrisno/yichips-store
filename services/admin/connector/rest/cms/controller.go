package cms

import (
	// native
	"strconv"

	// external
	"github.com/labstack/echo"

	// yichips
	"yichips/config"
	"yichips/exception"
	"yichips/utils/contract"
	"yichips/utils/rest"

	// admin service
	"yichips/services/admin"
	"yichips/services/admin/connector/rest/cms/request"
	"yichips/services/admin/data"
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
var mainService = admin.Main{}
var adminListLimit, _ = strconv.Atoi(config.App["admin_list_limit"])

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------
//

/**
 * Handle Admin Detail Show
 */
func (ctr *Controller) Index(e echo.Context) (err error) {
	// get page
	page, _ := strconv.Atoi(e.QueryParam("page"))

	// get admin details
	adminList, total, exc := mainService.FetchList(page, adminListLimit)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	data := map[string]contract.Model{
		"admins": adminList.WithMeta(page, adminListLimit, total),
	}

	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------

/**
 * Handle Admin Detail Show
 */
func (ctr *Controller) Show(e echo.Context) (err error) {
	// get path parameter
	adminCode := e.Param("code")

	// get admin details
	adminDetail, exc := mainService.GetDetail(adminCode)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	data := map[string]contract.Model{
		"admin_detail": adminDetail,
	}

	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------

/**
 * Handle New Admin Creation
 */
func (ctr *Controller) Create(e echo.Context) (err error) {
	// bind and validate request
	req := new(request.Create)
	e.Bind(req)
	if err = e.Validate(req); err != nil {
		return rest.ConstructErrorResponse(e, exception.NewInputValidationFailed(err.Error()))
	}

	// map req to input data
	reqData := data.NewNewAdminCreate(
		map[string]interface{}{
			"Name":     req.Name,
			"Email":    req.Email,
			"Password": req.Password,
		},
	)

	// get admin details
	adminDetail, exc := mainService.CreateNew(reqData)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	data := map[string]contract.Model{
		"created_admin": adminDetail,
	}

	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------

/**
 * Handle Existing Admin Update
 */
func (ctr *Controller) Update(e echo.Context) (err error) {
	// get path parameter
	adminCode := e.Param("code")

	// bind and validate request
	req := new(request.Update)
	e.Bind(req)
	if err = e.Validate(req); err != nil {
		return rest.ConstructErrorResponse(e, exception.NewInputValidationFailed(err.Error()))
	}

	// map req to input data
	reqData := data.NewExistingAdminUpdate(
		map[string]interface{}{
			"Name":     req.Name,
			"Email":    req.Email,
			"Password": req.Password,
			"Status":   req.Status,
		},
	)

	// get admin details
	adminDetail, exc := mainService.UpdateExisting(adminCode, reqData)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	data := map[string]contract.Model{
		"updated_admin": adminDetail,
	}

	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------

/**
 * Handle Admin Deletion
 */
func (ctr *Controller) Destroy(e echo.Context) (err error) {
	// get path parameter
	adminCode := e.Param("code")

	// get admin details
	deletedAdminCode, exc := mainService.DeleteExisting(adminCode)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	data := map[string]contract.Model{
		"deleted_admin": deletedAdminCode,
	}

	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------
