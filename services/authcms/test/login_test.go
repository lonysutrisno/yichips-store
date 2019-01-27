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
	"yichips/utils/helper/hash"

	// authcms service
	"yichips/services/authcms/output"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Test Login()
 */
func TestLogin(t *testing.T) {
	// ********************************************
	// ********************************************
	dPassword := "masuk123"
	dApiToken := "asd"
	expirationMinutes, _ := strconv.Atoi(os.Getenv("API_TOKEN_EXPIRATION_DURATION"))
	dApiExpiration := time.Now().Add(time.Minute * time.Duration(expirationMinutes))
	currentTime := dApiExpiration

	dummy := model.Admin{
		Code:               "ADM-123",
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           hash.Make(dPassword),
		Status:             "ACTIVE",
		ApiToken:           dApiToken,
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
	expectedExpiration := int(dApiExpiration.Unix())
	expectedCode := dummy.Code
	oldApiToken := dApiToken

	email := dummy.Email
	password := dPassword
	authorization, adminCode, exc := service.Login(email, password)

	assert.Equal(t, nil, exc, "[LOGIN - Success] exception should be nil")
	assert.Equal(t, expectedCode, adminCode, "[LOGIN - Success] admin code should be equal to expected admin code")
	assert.Equal(t, len(authorization.(output.AdminLogin).ApiToken), 44, "[LOGIN - Success] api token should be generated and hashed")
	assert.NotEqual(t, oldApiToken, authorization.(output.AdminLogin).ApiToken, "[LOGIN - Success] api token should NOT be equal to old api token")
	assert.Equal(t, expectedExpiration/10, authorization.(output.AdminLogin).ApiTokenExpiration/10, "[LOGIN - Success] object should have now + 2 hours new expiration")

	// ----------------------------------------
	// ----------------------------------------
	// 2) email not found test
	// ----------------------------------------
	// ----------------------------------------
	email = "random@email.com"
	password = dPassword
	authorization, adminCode, exc = service.Login(email, password)

	assert.Equal(t, nil, authorization, "[LOGIN - Email not found] object should be nil")
	assert.Equal(t, "", adminCode, "[LOGIN - Email not found] admin code should be empty string")
	assert.Equal(t, exception.NewLoginAuthenticationFailed(), exc, "[LOGIN - Email not found] exception should be LoginAuthenticationFailed")

	// ----------------------------------------
	// ----------------------------------------
	// 3) incorrect password test
	// ----------------------------------------
	// ----------------------------------------
	email = dummy.Email
	password = "random-password"
	authorization, adminCode, exc = service.Login(email, password)

	assert.Equal(t, nil, authorization, "[LOGIN - Incorrect password] object should be nil")
	assert.Equal(t, "", adminCode, "[LOGIN - Incorrect password] admin code should be empty string")
	assert.Equal(t, exception.NewLoginAuthenticationFailed(), exc, "[LOGIN - Incorrect password] exception should be LoginAuthenticationFailed")

	// ********************************************
	// ********************************************
	dummy.Status = "INACTIVE"

	DB.Save(&dummy)
	// ********************************************
	// ********************************************

	// ----------------------------------------
	// ----------------------------------------
	// 4) account inactive test
	// ----------------------------------------
	// ----------------------------------------
	email = dummy.Email
	password = dPassword
	authorization, adminCode, exc = service.Login(email, password)

	assert.Equal(t, nil, authorization, "[LOGIN - Account inactive] object should be nil")
	assert.Equal(t, "", adminCode, "[LOGIN - Account inactive] admin code should be empty string")
	assert.Equal(t, exception.NewAccountInactive(), exc, "[LOGIN - Account inactive] exception should be AccountInactive")

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}

/**
 * Benchmark Login()
 */
func BenchmarkLogin(b *testing.B) {
	// ********************************************
	// ********************************************
	dPassword := "masuk123"
	dApiToken := "asd"

	expirationMinutes, _ := strconv.Atoi(os.Getenv("API_TOKEN_EXPIRATION_DURATION"))
	currentTime := time.Now().Add(time.Minute * time.Duration(expirationMinutes))

	dummy := model.Admin{
		Code:               "ADM-123",
		Name:               "Test",
		Email:              "test@yichips.com",
		Password:           hash.Make(dPassword),
		Status:             "ACTIVE",
		ApiToken:           dApiToken,
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
	email := dummy.Email
	password := dPassword

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		service.Login(email, password)
	}
	b.StopTimer()

	// ********************************************
	// ********************************************
	DB.DropTable(&model.Admin{})
	// ********************************************
	// ********************************************
}
