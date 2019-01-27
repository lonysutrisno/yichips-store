package authcms

import (
	// native
	"reflect"

	// yichips
	"yichips/exception"
	"yichips/utils/contract"
	"yichips/utils/helper/hash"

	// authcms service
	"yichips/services/authcms/model"
	"yichips/services/authcms/output"
)

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Admin Auth Service
 |	------------------
 |
 |	This service is used to handle admin (cms) auth operations such as
 |	login, forgot password, etc.
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------
// ------------------------------------------------------

var AdminModel = model.NewAdmin()

// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Main struct {
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Admin Login
 *
 * @param 	string 	email
 * @param 	string 	password
 *
 * @return 	contract.ServiceOutput
 * @return 	string
 * @return 	exception.Exception
 */
func (main *Main) Login(email string, password string) (contract.ServiceOutput, string, exception.Exception) {
	// get admin data by login credential (email)
	admin, exc := AdminModel.GetByEmail(email)
	if exc != nil {
		// change exception if exception == record not found
		if reflect.TypeOf(exc) == reflect.TypeOf(exception.RecordNotFound{}) {
			return nil, "", exception.NewLoginAuthenticationFailed()
		}

		return nil, "", exc
	}

	// check hashed password
	if !hash.Check(password, admin.Password) {
		return nil, "", exception.NewLoginAuthenticationFailed()
	}

	// check if admin status not active
	if admin.Status != "ACTIVE" {
		return nil, "", exception.NewAccountInactive()
	}

	// update token
	admin.UpdateToken()

	return output.NewAdminLogin(admin), admin.Code, nil
}

/**
 * Admin Authorize
 *
 * @param 	string 	token
 *
 * @return 	contract.ServiceOutput
 * @return 	exception.Exception
 */
func (main *Main) Authorize(token string) (contract.ServiceOutput, exception.Exception) {
	// get admin data by api credential (token)
	admin, exc := AdminModel.GetByToken(token)
	if exc != nil {
		return nil, exception.NewAuthorizationFailed()
	}

	return output.NewAdminAuthorize(admin), nil
}
