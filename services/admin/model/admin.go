package model

import (
	// external
	"github.com/jinzhu/copier"

	// yichips
	"yichips/bootstrap"
	base "yichips/core/model"
	"yichips/exception"

	// admin service
	adminData "yichips/services/admin/data"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Admin struct {
	base.Admin `mapstructure:",squash"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------

var DB = bootstrap.DB

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Constructor
 *
 * @return 	Admin
 */
func NewAdmin() Admin {
	return Admin{}
}

// ------------------------------------------------------

/**
 * Fetch Admin Records
 *
 * @return 	[]Admin
 * @return 	exception.Exception
 */
func (admin *Admin) Fetch(offset int, limit int) ([]Admin, exception.Exception) {
	var adminList []Admin
	DB.Offset(offset).Limit(limit).Find(&adminList)

	return adminList, nil
}

// ------------------------------------------------------

/**
 * Count Total Admin Records
 *
 * @return 	int
 * @return 	exception.Exception
 */
func (admin *Admin) Count() (int, exception.Exception) {
	var adminList []Admin
	var total int
	DB.Find(&adminList).Count(&total)

	return total, nil
}

// ------------------------------------------------------

/**
 * Get Admin Detail by Code
 *
 * @param 	string code
 *
 * @return 	*Admin
 * @return 	exception.Exception
 */
func (admin *Admin) Get(code string) (*Admin, exception.Exception) {
	// retrieve admin detail
	var adminDetail Admin
	query := DB.Where(
		map[string]interface{}{
			"code": code,
		},
	).First(&adminDetail)

	// if record not found
	if query.RecordNotFound() {
		return nil, exception.NewRecordNotFound("Admin")
	}

	return &adminDetail, nil
}

// ------------------------------------------------------

/**
 * Store New Admin Record
 *
 * @param 	adminData.NewAdminCreate 	data
 *
 * @return 	*Admin
 * @return 	exception.Exception
 */
func (admin *Admin) Create(data adminData.NewAdminCreate) (*Admin, exception.Exception) {
	// check if email already has been used
	var existingAdmin Admin
	query := DB.Where(
		map[string]interface{}{
			"email": data.Email,
		},
	).First(&existingAdmin)
	if !query.RecordNotFound() {
		return nil, exception.NewDuplicateRecordFound("Email")
	}

	// store admin data
	var newAdmin Admin
	copier.Copy(&newAdmin, &data)
	DB.Create(&newAdmin)

	return &newAdmin, nil
}

// ------------------------------------------------------

/**
 * Update Existing Admin Record
 *
 * @param 	string 	code
 * @param 	map[string]interface{} 	data
 *
 * @return 	*Admin
 * @return 	exception.Exception
 */
func (admin *Admin) Update(code string, data adminData.ExistingAdminUpdate) (*Admin, exception.Exception) {
	// retrieve existing admin
	existingAdminDetail, exc := admin.Get(code)
	if exc != nil {
		return nil, exc
	}

	// check if email already has been used
	var existingAdmin Admin
	query := DB.Where(
		map[string]interface{}{
			"email": data.Email,
		},
	).Not(
		map[string]interface{}{
			"code": code,
		}).First(&existingAdmin)
	if !query.RecordNotFound() {
		return nil, exception.NewDuplicateRecordFound("Email")
	}

	// update admin data
	DB.Model(&existingAdminDetail).Updates(data)

	return existingAdminDetail, nil
}

// ------------------------------------------------------

/**
 * Delete Existing Admin Record
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
func (admin *Admin) Delete(code string) (string, exception.Exception) {
	// check if code exists
	var existingAdmin Admin
	query := DB.Where(
		map[string]interface{}{
			"code": code,
		},
	).First(&existingAdmin)
	if query.RecordNotFound() {
		return "", exception.NewRecordNotFound("Admin")
	}

	// get admin code
	adminCode := existingAdmin.Code

	// delete admin
	DB.Delete(existingAdmin)

	return adminCode, nil
}
