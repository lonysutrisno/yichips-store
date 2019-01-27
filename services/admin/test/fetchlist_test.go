package test

import (
	// native
	"testing"
	"time"

	// external
	"github.com/stretchr/testify/assert"

	// yichips
	"yichips/core/model"
	"yichips/utils/collection"

	// admin service
	"yichips/services/admin/output"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Test FetchList()
 */
func TestFetchList(t *testing.T) {
	// ********************************************
	// ********************************************
	currentTime := time.Now()
	dPage := 1
	dLimit := 15

	dummy1 := model.Admin{
		Code:               "ADM-123",
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           "masuk123",
		Status:             "ACTIVE",
		ApiToken:           "asd",
		ApiTokenExpiration: &currentTime,
	}
	dummy2 := model.Admin{
		Code:               "ADM-456",
		Name:               "Test2",
		Email:              "test2@yichips.com",
		Password:           "qwerty123",
		Status:             "INACTIVE",
		ApiToken:           "qwe",
		ApiTokenExpiration: &currentTime,
	}

	if DB.HasTable(&model.Admin{}) {
		DB.DropTable(&model.Admin{})
	}
	DB.CreateTable(&model.Admin{})
	DB.Create(&dummy1)
	DB.Create(&dummy2)
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// 1) success test
	// ----------------------------------------
	// ----------------------------------------
	expectedSuccess1 := output.AdminDetail{
		Code:   dummy1.Code,
		Name:   dummy1.Name,
		Email:  dummy1.Email,
		Status: dummy1.Status,
	}

	expectedSuccess2 := output.AdminDetail{
		Code:   dummy2.Code,
		Name:   dummy2.Name,
		Email:  dummy2.Email,
		Status: dummy2.Status,
	}

	expectedTotalSuccessRecord := 2

	expectedListSuccess := collection.NewModelResponse()
	expectedListSuccess.Add(expectedSuccess1)
	expectedListSuccess.Add(expectedSuccess2)

	page := dPage
	limit := dLimit

	adminList, totalRecord, exc := service.FetchList(page, limit)

	assert.Equal(t, nil, exc, "[GET DETAIL - Success] exception should be nil")
	assert.Equal(t, expectedListSuccess, adminList, "[GET DETAIL - Success] object list should be equal to expected success list object")
	assert.Equal(t, expectedTotalSuccessRecord, totalRecord, "[GET DETAIL - Success] total record should be equal to expected total success record")

	// ----------------------------------------
	// ----------------------------------------
	// 2) limited record return success
	// ----------------------------------------
	// ----------------------------------------
	expectedSuccess := output.AdminDetail{
		Code:   dummy1.Code,
		Name:   dummy1.Name,
		Email:  dummy1.Email,
		Status: dummy1.Status,
	}

	expectedTotalSuccessRecord = 2

	expectedListSuccess = collection.NewModelResponse()
	expectedListSuccess.Add(expectedSuccess)

	page = dPage
	limit = 1

	adminList, totalRecord, exc = service.FetchList(page, limit)

	assert.Equal(t, nil, exc, "[GET DETAIL - Limited record return success] exception should be nil")
	assert.Equal(t, expectedListSuccess, adminList, "[GET DETAIL - Limited record return success] object list should be equal to expected success list object")
	assert.Equal(t, expectedTotalSuccessRecord, totalRecord, "[GET DETAIL - Limited record return success] total record should be equal to expected total success record")

	// ********************************************
	// ********************************************
	DB.Delete(&dummy1)
	DB.Delete(&dummy2)
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// 3) empty success test
	// ----------------------------------------
	// ----------------------------------------
	adminList, totalRecord, exc = service.FetchList(page, limit)

	assert.Equal(t, nil, exc, "[GET DETAIL - Empty success] exception should be nil")
	assert.Equal(t, collection.NewModelResponse(), adminList, "[GET DETAIL - Empty success] object list should be empty list")
	assert.Equal(t, 0, totalRecord, "[GET DETAIL - Success] total record should be zero")

	// ----------------------------------------

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}

/**
 * Benchmark FetchList()
 */
func BenchmarkFetchList(b *testing.B) {
	// ********************************************
	// ********************************************
	currentTime := time.Now()
	dPage := 1
	dLimit := 15

	dummy1 := model.Admin{
		Code:               "ADM-123",
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           "masuk123",
		Status:             "ACTIVE",
		ApiToken:           "asd",
		ApiTokenExpiration: &currentTime,
	}
	dummy2 := model.Admin{
		Code:               "ADM-456",
		Name:               "Test2",
		Email:              "test2@yichips.com",
		Password:           "qwerty123",
		Status:             "INACTIVE",
		ApiToken:           "qwe",
		ApiTokenExpiration: &currentTime,
	}

	if DB.HasTable(&model.Admin{}) {
		DB.DropTable(&model.Admin{})
	}
	DB.CreateTable(&model.Admin{})
	DB.Create(&dummy1)
	DB.Create(&dummy2)
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// benchmark
	// ----------------------------------------
	// ----------------------------------------
	page := dPage
	limit := dLimit

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		service.FetchList(page, limit)
	}
	b.StopTimer()

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}
