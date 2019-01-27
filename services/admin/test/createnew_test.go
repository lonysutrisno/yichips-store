package test

import (
	// native
	"testing"

	// external
	"github.com/stretchr/testify/assert"

	// yichips
	"yichips/core/model"
	"yichips/exception"
	"yichips/utils/helper/hash"

	// admin service
	"yichips/services/admin/data"
	"yichips/services/admin/output"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Test CreateNew()
 */
func TestCreateNew(t *testing.T) {
	// ********************************************
	// ********************************************
	dName := "Test"
	dEmail := "test@yichips.com"
	dPassword := "masuk123"

	newData := data.NewNewAdminCreate(
		map[string]interface{}{
			"Name":     dName,
			"Email":    dEmail,
			"Password": dPassword,
		},
	)

	dummy := model.Admin{
		Name:   dName,
		Email:  dEmail,
		Status: "ACTIVE",
	}

	if DB.HasTable(&model.Admin{}) {
		DB.DropTable(&model.Admin{})
	}
	DB.CreateTable(&model.Admin{})
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// 1) success test
	// ----------------------------------------
	// ----------------------------------------
	adminDetail, exc := service.CreateNew(newData)
	normalizedAdminDetail := adminDetail.(output.AdminDetail)

	assert.Equal(t, nil, exc, "[CREATE NEW - Success] exception should be nil")
	assert.NotEqual(t, 0, len(normalizedAdminDetail.Code), "[CREATE NEW - Success] expected success object code should not be empty")
	assert.Equal(t, dummy.Name, normalizedAdminDetail.Name, "[CREATE NEW - Success] inserted name should be equal to expected success object name")
	assert.Equal(t, dummy.Email, normalizedAdminDetail.Email, "[CREATE NEW - Success] inserted email should be equal to expected success object email")
	assert.Equal(t, dummy.Status, normalizedAdminDetail.Status, "[CREATE NEW - Success] inserted status should be equal to expected success object status")

	// ********************************************
	// ********************************************
	expectedDBData := model.Admin{}
	query := DB.Where(
		map[string]interface{}{
			"code": normalizedAdminDetail.Code,
		},
	).First(&expectedDBData)
	// ********************************************
	// ********************************************

	assert.Equal(t, false, query.RecordNotFound(), "[CREATE NEW - Success] data should be saved to DB")
	assert.Equal(t, hash.Make(dPassword), expectedDBData.Password, "[CREATE NEW - Success] password saved in DB should be hashed")

	// ----------------------------------------
	// ----------------------------------------
	// 2) duplicate email found test
	// ----------------------------------------
	// ----------------------------------------
	newData.Name = "random for duplicate"
	newData.Email = dEmail

	adminDetail, exc = service.CreateNew(newData)

	assert.Equal(t, nil, adminDetail, "[CREATE NEW - Duplicate email found] adminDetail should be nil")
	assert.Equal(t, exception.NewDuplicateRecordFound("Email"), exc, "[CREATE NEW - Duplicate email found] exception should be DuplicateRecordFound")

	// ********************************************
	// ********************************************
	expectedDBData = model.Admin{}
	query = DB.Where(
		map[string]interface{}{
			"name":  newData.Name,
			"email": newData.Email,
		},
	).First(&expectedDBData)
	// ********************************************
	// ********************************************

	assert.Equal(t, true, query.RecordNotFound(), "[CREATE NEW - Duplicate email found] data should NOT be saved to DB")

	// ********************************************
	// ********************************************
	DB.Where(
		map[string]interface{}{
			"email": dEmail,
		},
	).Delete(&model.Admin{})
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// 3) Soft-deleted email reuse success
	// ----------------------------------------
	// ----------------------------------------
	adminDetail, exc = service.CreateNew(newData)
	normalizedAdminDetail = adminDetail.(output.AdminDetail)

	assert.Equal(t, nil, exc, "[CREATE NEW - Soft-deleted email reuse success] exception should be nil")
	assert.Equal(t, dummy.Email, normalizedAdminDetail.Email, "[CREATE NEW - Soft-deleted email reuse success] inserted email should be equal to expected success object email")

	// ********************************************
	// ********************************************
	expectedDBData = model.Admin{}
	query = DB.Where(
		map[string]interface{}{
			"email": newData.Email,
		},
	).First(&expectedDBData)
	// ********************************************
	// ********************************************

	assert.Equal(t, false, query.RecordNotFound(), "[CREATE NEW - Soft-deleted email reuse success] data should be saved to DB")

	// ----------------------------------------

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}

/**
 * Benchmark CreateNew()
 */
func BenchmarkCreateNew(b *testing.B) {
	// ********************************************
	// ********************************************
	dEmail := "test@yichips.com"

	if DB.HasTable(&model.Admin{}) {
		DB.DropTable(&model.Admin{})
	}
	DB.CreateTable(&model.Admin{})
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// benchmark
	// ----------------------------------------
	// ----------------------------------------
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		newData := data.NewNewAdminCreate(
			map[string]interface{}{
				"Name":     "Test",
				"Email":    dEmail,
				"Password": "masuk123",
			},
		)

		service.CreateNew(newData)

		DB.Unscoped().Delete(&model.Admin{})
	}
	b.StopTimer()

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}
