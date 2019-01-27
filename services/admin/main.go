package admin

import (
	// native
	// "strings"

	// yichips
	"yichips/exception"
	"yichips/utils/collection"
	"yichips/utils/contract"

	// admin service
	dataRaw "yichips/services/admin/data"
	"yichips/services/admin/model"
	"yichips/services/admin/output"
)

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Admin Service
 |	------------------
 |
 |	This service is used to handle admin management operations
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------

var AdminModel = model.NewAdmin()

// ------------------------------------------------------
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
 * Fetch Admin List
 *
 * @param 	int 	page
 * @param 	int 	limit
 *
 * @return 	collection.ModelResponse
 * @return 	int
 * @return 	exception.Exception
 */
func (main *Main) FetchList(page int, limit int) (collection.ModelResponse, int, exception.Exception) {
	// get admin list
	admins, exc := AdminModel.Fetch(((page - 1) * limit), limit)
	if exc != nil {
		return collection.NewModelResponse(), 0, exc
	}

	// get total list
	total, exc := AdminModel.Count()
	if exc != nil {
		return collection.NewModelResponse(), 0, exc
	}

	return output.NewAdminDetailList(admins), total, nil
}

// ------------------------------------------------------

/**
 * Get Admin Detail
 *
 * @param 	string 	code
 *
 * @return 	contract.ServiceOutput
 * @return 	exception.Exception
 */
func (main *Main) GetDetail(code string) (contract.ServiceOutput, exception.Exception) {
	// get admin data by code
	admin, exc := AdminModel.Get(code)
	if exc != nil {
		return nil, exc
	}

	return output.NewAdminDetail(admin), nil
}

// ------------------------------------------------------

/**
 * Create New Admin
 *
 * @param 	dataRaw.NewAdminCreate 	data
 *
 * @return 	contract.ServiceOutput
 * @return 	exception.Exception
 */
func (main *Main) CreateNew(data dataRaw.NewAdminCreate) (contract.ServiceOutput, exception.Exception) {
	// create admin
	admin, exc := AdminModel.Create(data)
	if exc != nil {
		return nil, exc
	}

	return output.NewAdminDetail(admin), nil
}

// ------------------------------------------------------

/**
 * Create New Admin
 *
 * @param 	string 	code
 * @param 	dataRaw.ExistingAdminUpdate 	data
 *
 * @return 	contract.ServiceOutput
 * @return 	exception.Exception
 */
func (main *Main) UpdateExisting(code string, data dataRaw.ExistingAdminUpdate) (contract.ServiceOutput, exception.Exception) {
	// update admin
	admin, exc := AdminModel.Update(code, data)
	if exc != nil {
		return nil, exc
	}

	return output.NewAdminDetail(admin), nil
}

// ------------------------------------------------------

/**
 * Delete Existing Admin
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
func (main *Main) DeleteExisting(code string) (string, exception.Exception) {
	// delete admin data by code
	adminCode, exc := AdminModel.Delete(code)
	if exc != nil {
		return "", exc
	}

	return adminCode, nil
}

// ------------------------------------------------------
