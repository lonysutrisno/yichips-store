package test

import (
	// native
	"os"
	"strconv"
	"testing"
	"time"

	// external
	"github.com/stretchr/testify/assert"

	// yichips
	"yichips/core/model"
	"yichips/exception"

	// authcms service
	"yichips/services/authcms/output"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Test Authorize()
 */
func TestAuthorize(t *testing.T) {
	// ********************************************
	// ********************************************
	expirationMinutes, _ := strconv.Atoi(os.Getenv("API_TOKEN_EXPIRATION_DURATION"))
	dCurrentTime := time.Now().Add(time.Minute * time.Duration(expirationMinutes))
	dApiToken := "asd"
	dStatus := "ACTIVE"

	dummy := model.Admin{
		Code:               "ADM-123",
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           "masuk123",
		Status:             dStatus,
		ApiToken:           dApiToken,
		ApiTokenExpiration: &dCurrentTime,
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
	expectedSuccess := output.AdminAuthorize{
		ID:     1,
		Code:   dummy.Code,
		Name:   dummy.Name,
		Email:  dummy.Email,
		Status: dummy.Status,
	}

	expectedStatus := dStatus

	token := dApiToken
	adminDetail, exc := service.Authorize(token)

	assert.Equal(t, nil, exc, "[AUTHORIZE - Success] exception should be nil")
	assert.Equal(t, expectedSuccess, adminDetail.(output.AdminAuthorize), "[AUTHORIZE - Success] object should be equal to expected success object")
	assert.Equal(t, expectedStatus, adminDetail.(output.AdminAuthorize).Status, "[AUTHORIZE - Success] status should be 'ACTIVE'")

	// ----------------------------------------
	// ----------------------------------------
	// 2) token not found test
	// ----------------------------------------
	// ----------------------------------------
	token = "randomtoken"
	adminDetail, exc = service.Authorize(token)

	assert.Equal(t, nil, adminDetail, "[AUTHORIZE - Token not found] object should be nil")
	assert.Equal(t, exception.NewAuthorizationFailed(), exc, "[AUTHORIZE - Token not found] exception should be 	adminDetail, exc := service.Authorize(token)")

	// ********************************************
	// ********************************************
	expirationMinutes, _ = strconv.Atoi(os.Getenv("API_TOKEN_EXPIRATION_DURATION"))
	expiredTime := dCurrentTime.Add(time.Minute * time.Duration(-expirationMinutes-60))

	dummy.ApiTokenExpiration = &expiredTime

	DB.Save(&dummy)
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// 3) expired token test
	// ----------------------------------------
	// ----------------------------------------
	token = dApiToken
	adminDetail, exc = service.Authorize(token)

	assert.Equal(t, nil, adminDetail, "[AUTHORIZE - Api token expired] object should be nil")
	assert.Equal(t, exception.NewAuthorizationFailed(), exc, "[AUTHORIZE - Api token expired] exception should be AuthorizationFailed")

	// ********************************************
	// ********************************************
	dummy.Status = "INACTIVE"
	dummy.ApiTokenExpiration = &dCurrentTime

	DB.Save(&dummy)
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// 4) account inactive test
	// ----------------------------------------
	// ----------------------------------------
	token = dApiToken
	adminDetail, exc = service.Authorize(token)

	assert.Equal(t, nil, adminDetail, "[AUTHORIZE - Account inactive] object should be nil")
	assert.Equal(t, exception.NewAuthorizationFailed(), exc, "[AUTHORIZE - Account inactive] exception should be AuthorizationFailed")

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}

/**
 * Benchmark Authorize()
 */
func BenchmarkAuthorize(b *testing.B) {
	// ********************************************
	// ********************************************
	expirationMinutes, _ := strconv.Atoi(os.Getenv("API_TOKEN_EXPIRATION_DURATION"))
	dCurrentTime := time.Now().Add(time.Minute * time.Duration(expirationMinutes))
	dApiToken := "asd"
	dStatus := "ACTIVE"

	dummy := model.Admin{
		Code:               "ADM-123",
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           "masuk123",
		Status:             dStatus,
		ApiToken:           dApiToken,
		ApiTokenExpiration: &dCurrentTime,
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
	token := dApiToken

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		service.Authorize(token)
	}
	b.StopTimer()

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}
