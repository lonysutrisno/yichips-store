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

	// admin service
	"yichips/services/admin/output"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Test GetDetail()
 */
func TestGetDetail(t *testing.T) {
	// ********************************************
	// ********************************************
	currentTime := time.Now()

	dummy := model.Admin{
		Code:               "ADM-123",
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           "masuk123",
		Status:             "ACTIVE",
		ApiToken:           "asd",
		ApiTokenExpiration: &currentTime,
	}

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
		Code:   dummy.Code,
		Name:   dummy.Name,
		Email:  dummy.Email,
		Status: dummy.Status,
	}

	code := dummy.Code
	adminDetail, exc := service.GetDetail(code)

	assert.Equal(t, nil, exc, "[GET DETAIL - Success] exception should be nil")
	assert.Equal(t, expectedSuccess, adminDetail.(output.AdminDetail), "[GET DETAIL - Success] object should be equal to expected success object")

	// ----------------------------------------
	// ----------------------------------------
	// 2) admin code not found test
	// ----------------------------------------
	// ----------------------------------------
	code = "RAND-CODE"
	adminDetail, exc = service.GetDetail(code)

	assert.Equal(t, nil, adminDetail, "[GET DETAIL - Admin code not found] object should be nil")
	assert.Equal(t, exception.NewRecordNotFound("Admin"), exc, "[GET DETAIL - Admin code not found] exception should be RecordNotFound")

	// ----------------------------------------

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}

/**
 * Benchmark GetDetail()
 */
func BenchmarkGetDetail(b *testing.B) {
	// ********************************************
	// ********************************************
	currentTime := time.Now()

	dummy := model.Admin{
		Code:               "ADM-123",
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           "masuk123",
		Status:             "ACTIVE",
		ApiToken:           "asd",
		ApiTokenExpiration: &currentTime,
	}

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
		service.GetDetail(code)
	}
	b.StopTimer()

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}
