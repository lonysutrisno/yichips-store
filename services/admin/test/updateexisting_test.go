package test

import (
	// native
	"testing"
	"time"

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
 * Test UpdateExisting()
 */
func TestUpdateExisting(t *testing.T) {
	// ********************************************
	// ********************************************
	currentTime := time.Now()

	dCode := "ADM-123"

	dummy := model.Admin{
		Code:               dCode,
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           "masuk123",
		Status:             "ACTIVE",
		ApiToken:           "asd",
		ApiTokenExpiration: &currentTime,
	}

	dName := "Test Update"
	dEmail := "email@update.com"
	dPassword := "password-update"
	dStatus := "INACTIVE"

	newData := data.NewExistingAdminUpdate(
		map[string]interface{}{
			"Name":     dName,
			"Email":    dEmail,
			"Password": dPassword,
			"Status":   dStatus,
		},
	)

	if DB.HasTable(&model.Admin{}) {
		DB.DropTable(&model.Admin{})
	}
	DB.CreateTable(&model.Admin{})
	DB.Create(&dummy)
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// 1) success test
	// ----------------------------------------
	// ----------------------------------------
	expectedSuccess := output.AdminDetail{
		Code:   dCode,
		Name:   dName,
		Email:  dEmail,
		Status: dStatus,
	}

	code := dummy.Code
	adminDetail, exc := service.UpdateExisting(code, newData)

	assert.Equal(t, nil, exc, "[UPDATE EXISTING - Success] exception should be nil")
	assert.Equal(t, expectedSuccess, adminDetail.(output.AdminDetail), "[UPDATE EXISTING - Success] object should be equal to expected success object")

	// ----------------------------------------
	// ----------------------------------------
	// 2) update with old data success test
	// ----------------------------------------
	// ----------------------------------------
	expectedSuccess = output.AdminDetail{
		Code:   dummy.Code,
		Name:   newData.Name,
		Email:  newData.Email,
		Status: newData.Status,
	}

	code = dummy.Code
	adminDetail, exc = service.UpdateExisting(code, newData)

	assert.Equal(t, nil, exc, "[UPDATE EXISTING - Update with old data] exception should be nil")
	assert.Equal(t, expectedSuccess, adminDetail.(output.AdminDetail), "[UPDATE EXISTING - Update with old data] object should be equal to old expected success object")

	// ********************************************
	// ********************************************
	expectedDBData := model.Admin{}
	DB.Where(
		map[string]interface{}{
			"code": code,
		},
	).First(&expectedDBData)
	// ********************************************
	// ********************************************

	assert.Equal(t, hash.Make(dPassword), expectedDBData.Password, "[CREATE NEW - Success] password saved in DB should be hashed")

	// ----------------------------------------
	// ----------------------------------------
	// 3) admin code not found test
	// ----------------------------------------
	// ----------------------------------------
	code = "RAND-CODE"
	adminDetail, exc = service.UpdateExisting(code, newData)

	assert.Equal(t, nil, adminDetail, "[UPDATE EXISTING - Admin code not found] object should be nil")
	assert.Equal(t, exception.NewRecordNotFound("Admin"), exc, "[UPDATE EXISTING - Admin code not found] exception should be RecordNotFound")

	// ********************************************
	// ********************************************
	dDuplicateEmail := "testForDuplicate@yichips.com"
	dDuplicateName := "Doppleganger Test"
	dDuplicateCode := "ADM-234"

	anotherDummy := model.Admin{
		Code:               dDuplicateCode,
		Name:               dDuplicateName,
		Email:              dDuplicateEmail,
		Password:           "masuk123",
		Status:             "ACTIVE",
		ApiToken:           "asd",
		ApiTokenExpiration: &currentTime,
	}

	DB.Create(&anotherDummy)
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// 4) duplicate email found test
	// ----------------------------------------
	// ----------------------------------------
	code = dCode
	newData.Email = dDuplicateEmail

	adminDetail, exc = service.UpdateExisting(code, newData)

	assert.Equal(t, nil, adminDetail, "[UPDATE EXISTING - Duplicate email found] adminDetail should be nil")
	assert.Equal(t, exception.NewDuplicateRecordFound("Email"), exc, "[UPDATE EXISTING - Duplicate email found] exception should be DuplicateRecordFound")

	// ********************************************
	// ********************************************
	expectedDBData = model.Admin{}
	query := DB.Where(
		map[string]interface{}{
			"name":  newData.Name,
			"email": newData.Email,
		},
	).First(&expectedDBData)
	// ********************************************
	// ********************************************

	assert.Equal(t, true, query.RecordNotFound(), "[CREATE NEW - Duplicate email found] data should NOT be Changed to DB")

	// ----------------------------------------

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}

/**
 * Benchmark UpdateExisting()
 */
func BenchmarkUpdateExisting(b *testing.B) {
	// ********************************************
	// ********************************************
	currentTime := time.Now()

	dCode := "ADM-123"

	dummy := model.Admin{
		Code:               dCode,
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           "masuk123",
		Status:             "ACTIVE",
		ApiToken:           "asd",
		ApiTokenExpiration: &currentTime,
	}

	dName := "Test Update"
	dEmail := "eamil@update.com"
	dPassword := "password-update"
	dStatus := "INACTIVE"

	newData := data.NewExistingAdminUpdate(
		map[string]interface{}{
			"Name":     dName,
			"Email":    dEmail,
			"Password": dPassword,
			"Status":   dStatus,
		},
	)

	if DB.HasTable(&model.Admin{}) {
		DB.DropTable(&model.Admin{})
	}
	DB.CreateTable(&model.Admin{})
	DB.Create(&dummy)
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// benchmark
	// ----------------------------------------
	// ----------------------------------------
	code := dummy.Code

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		service.UpdateExisting(code, newData)
	}
	b.StopTimer()

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}
