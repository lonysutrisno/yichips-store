package user

import (
	"encoding/base64"
	"fmt"
	"yichips/utils/helper/mail"

	// native
	// "strings"

	// yichips
	"reflect"
	"strconv"
	"yichips/config"
	"yichips/exception"
	"yichips/utils/collection"
	"yichips/utils/contract"
	"yichips/utils/helper/hash"
	"yichips/utils/rest"

	// user service
	input "yichips/services/user/input/form"
	request "yichips/services/user/input/requestValidator"
	"yichips/services/user/model"
	"yichips/services/user/output"

	"github.com/labstack/echo"
)

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	User Service
 |	------------------
 |
 |	This service is used to handle user management operations
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------

var UserModel = model.NewUser()

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Main struct {
}

var userListLimit, _ = strconv.Atoi(config.App["user_list_limit"])

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Handle User Register
 */
func (main *Main) Register(e echo.Context) (err error) {

	// get request and validate
	req := new(request.Create)
	e.Bind(req)
	if err = e.Validate(req); err != nil {
		return rest.ConstructErrorResponse(e, exception.NewInputValidationFailed(err.Error()))
	}
	// map req to input data
	reqData := input.NewNewUserCreate(
		map[string]interface{}{
			"Name":     req.Name,
			"Email":    req.Email,
			"Password": req.Password,
			"Phone":    req.Phone,
		},
	)
	//insert data to db
	user, exc := UserModel.Create(reqData)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}
	data := map[string]contract.Model{
		"created_user": user,
	}
	mail.SendMail(req.Email, "Email Verification", mail.CreateLink(user.Code, e))
	return rest.ConstructSuccessResponse(e, data)
}

/**
 * Handle User Login
 */
func (main *Main) LoginByMail(e echo.Context) (err error) {
	req := new(request.Login)
	e.Bind(req)
	if err = e.Validate(req); err != nil {
		return rest.ConstructErrorResponse(e, exception.NewInputValidationFailed(err.Error()))
	}

	// attempt to login user

	user, exc := UserModel.GetByEmail(req.Email)
	if exc != nil {
		// change exception if exception == record not found
		if reflect.TypeOf(exc) == reflect.TypeOf(exception.RecordNotFound{}) {
			return rest.ConstructErrorResponse(e, exception.NewLoginAuthenticationFailed())
		}

		return rest.ConstructErrorResponse(e, exc)
	}

	// check hashed password
	if !hash.Check(req.Password, user.Password) {
		return rest.ConstructErrorResponse(e, exception.NewLoginAuthenticationFailed())
	}

	if user.Status != "ACTIVE" {
		return rest.ConstructErrorResponse(e, exception.NewAccountInactive())

	}

	// get user details

	userDetail, exc := UserModel.Get("code", user.Code)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// return output.NewUserDetail(user), nil
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}
	user.UpdateToken()
	// prepare data
	data := map[string]contract.Model{
		"authorization": output.NewUserLogin(user),
		"user":          output.NewUserDetail(userDetail),
	}

	return rest.ConstructSuccessResponse(e, data)

}

/**
 * Handle User Detail Show
 */
func (main *Main) GetDetail(e echo.Context) (err error) {
	// get path parameter
	userCode := e.Param("code")

	// get user details
	userDetail, exc := UserModel.Get("code", userCode)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	data := map[string]contract.Model{
		"user_detail": output.NewUserDetail(userDetail),
	}

	return rest.ConstructSuccessResponse(e, data)
}

/**
 * Handle User Detail Show
 */
func (main *Main) MailVerification(e echo.Context) (err error) {
	// get path parameter
	token := e.Param("token")
	userCode, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// get user details
	_, exc := UserModel.Get("code", string(userCode))
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	userDetail, exc := UserModel.Update(string(userCode), map[string]interface{}{
		"Status": "ACTIVE",
	})
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}
	// prepare data
	data := map[string]contract.Model{
		"user_detail": output.NewUserDetail(userDetail),
	}

	return rest.ConstructSuccessResponse(e, data)
}

/**
 * Fetch User List
 *
 * @param 	int 	page
 * @param 	int 	limit
 *
 * @return 	collection.ModelResponse
 * @return 	int
 * @return 	exception.Exception
 */
func (main *Main) FetchList(page int, limit int) (collection.ModelResponse, int, exception.Exception) {
	// get user list
	users, exc := UserModel.Fetch(((page - 1) * limit), limit)
	if exc != nil {
		return collection.NewModelResponse(), 0, exc
	}

	// get total list
	total, exc := UserModel.Count()
	if exc != nil {
		return collection.NewModelResponse(), 0, exc
	}

	return output.NewUserDetailList(users), total, nil
}

// ------------------------------------------------------

/**
 * Create New User
 *
 * @param 	dataRaw.NewUserCreate 	data
 *
 * @return 	contract.ServiceOutput
 * @return 	exception.Exception
 */
// func (main *Main) CreateNew(data dataRaw.NewUserCreate) (contract.ServiceOutput, exception.Exception) {
// 	// create user
// 	user, exc := UserModel.Create(data)
// 	if exc != nil {
// 		return nil, exc
// 	}

// 	return output.NewUserDetail(user), nil
// }

// ------------------------------------------------------

/**
 * Delete Existing User
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
// func (main *Main) DeleteExisting(code string) (string, exception.Exception) {
// 	// delete user data by code
// 	userCode, exc := UserModel.Delete(code)
// 	if exc != nil {
// 		return "", exc
// 	}

// 	return userCode, nil
// }

// ------------------------------------------------------
