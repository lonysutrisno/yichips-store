package model

import (
	// native
	"os"
	"strconv"
	"time"

	// yichips
	"yichips/bootstrap"
	base "yichips/core/model"
	"yichips/exception"
	"yichips/utils/helper/hash"
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
 * Get Admin Detail by Email
 *
 * @param 	string email
 *
 * @return 	*Admin
 * @return 	exception.Exception
 */
func (admin *Admin) GetByEmail(email string) (*Admin, exception.Exception) {
	// retrieve admin detail
	var adminDetail Admin
	query := DB.Where(
		map[string]interface{}{
			"email": email,
		},
	).First(&adminDetail)

	// if record not found
	if query.RecordNotFound() {
		return nil, exception.NewRecordNotFound("")
	}

	return &adminDetail, nil
}

// ------------------------------------------------------

/**
 * Get Admin Detail by Token
 *
 * @param 	string email
 *
 * @return 	*Admin
 * @return 	exception.Exception
 */
func (admin *Admin) GetByToken(token string) (*Admin, exception.Exception) {
	// retrieve admin detail
	var adminDetail Admin
	query := DB.Where(
		map[string]interface{}{
			"api_token": token,
			"status":    "ACTIVE",
		},
	).Where("api_token_expiration > ?", time.Now()).First(&adminDetail)

	// if record not found
	if query.RecordNotFound() {
		return nil, exception.NewRecordNotFound("")
	}

	return &adminDetail, nil
}

// ------------------------------------------------------

/**
 * Update Admin Token
 */
func (admin *Admin) UpdateToken() {
	// update api token and the expiration
	expirationMinutes, _ := strconv.Atoi(os.Getenv("API_TOKEN_EXPIRATION_DURATION"))
	DB.Model(&admin).Updates(
		map[string]interface{}{
			"api_token":            hash.Randomize(),
			"api_token_expiration": time.Now().Add(time.Minute * time.Duration(expirationMinutes)),
		},
	)
}
