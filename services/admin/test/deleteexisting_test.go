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
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Test DeleteExisting()
 */
func TestDeleteExisting(t *testing.T) {
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
	code := dummy.Code
	adminCode, exc := service.DeleteExisting(code)

	assert.Equal(t, nil, exc, "[DELETE EXISTING - Success] exception should be nil")
	assert.Equal(t, code, adminCode, "[DELETE EXISTING - Success] deleted admin code should equal to inputted code")

	// ----------------------------------------
	// ----------------------------------------
	// 2) admin code not found test
	// ----------------------------------------
	// ----------------------------------------
	code = "RAND-CODE"
	adminCode, exc = service.DeleteExisting(code)

	assert.Equal(t, "", adminCode, "[DELETE EXISTING - Admin code not found] returned admin code should be empty")
	assert.Equal(t, exception.NewRecordNotFound("Admin"), exc, "[DELETE EXISTING - Admin code not found] exception should be RecordNotFound")

	// ----------------------------------------

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}

/**
 * Benchmark DeleteExisting()
 */
func BenchmarkDeleteExisting(b *testing.B) {
	// ********************************************
	// ********************************************
	currentTime := time.Now()

	code := "ADM-123"
	var dummy model.Admin

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
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dummy = model.Admin{
			Code:               code,
			Name:               "Test",
			Email:              "test@yichips.com",
			Password:           "masuk123",
			Status:             "ACTIVE",
			ApiToken:           "asd",
			ApiTokenExpiration: &currentTime,
		}
		DB.Create(&dummy)

		service.DeleteExisting(code)
	}
	b.StopTimer()

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}
