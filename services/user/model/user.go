package model

import (
	// external
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/copier"

	// yichips
	"yichips/bootstrap"
	base "yichips/core/model"
	"yichips/exception"
	"yichips/utils/helper/hash"

	// user service
	userData "yichips/services/user/input/form"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type User struct {
	base.User `mapstructure:",squash"`
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
 * @return 	User
 */
func NewUser() User {
	return User{}
}

// ------------------------------------------------------

/**
 * Fetch User Records
 *
 * @return 	[]User
 * @return 	exception.Exception
 */
func (user *User) Fetch(offset int, limit int) ([]User, exception.Exception) {
	var userList []User
	DB.Offset(offset).Limit(limit).Find(&userList)

	return userList, nil
}

// ------------------------------------------------------

/**
 * Count Total User Records
 *
 * @return 	int
 * @return 	exception.Exception
 */
func (user *User) Count() (int, exception.Exception) {
	var userList []User
	var total int
	DB.Find(&userList).Count(&total)

	return total, nil
}

// ------------------------------------------------------

/**
 * Get User Detail by Code
 *
 * @param 	string code
 *
 * @return 	*User
 * @return 	exception.Exception
 */
func (user *User) Get(field, code string) (*User, exception.Exception) {
	// retrieve user detail
	var userDetail User
	query := DB.Where(
		map[string]interface{}{
			field: code,
		},
	).First(&userDetail)

	// if record not found
	if query.RecordNotFound() {
		return nil, exception.NewRecordNotFound("User")
	}

	return &userDetail, nil
}

// ------------------------------------------------------

/**
 * Store New User Record
 *
 * @param 	userData.NewUserCreate 	data
 *
 * @return 	*User
 * @return 	exception.Exception
 */
func (user *User) Create(data userData.NewUserCreate) (*User, exception.Exception) {
	// check if email already has been used
	var existingUser User
	query := DB.Where(
		map[string]interface{}{
			"email": data.Email,
		},
	).First(&existingUser)
	if !query.RecordNotFound() {
		return nil, exception.NewDuplicateRecordFound("Email")
	}

	// store user data
	var newUser User
	copier.Copy(&newUser, &data)
	DB.Create(&newUser)

	return &newUser, nil
}

// ------------------------------------------------------

/**
 * Update Existing User Record
 *
 * @param 	string 	code
 * @param 	map[string]interface{} 	data
 *
 * @return 	*User
 * @return 	exception.Exception
 */
func (user *User) Update(code string, data map[string]interface{}) (*User, exception.Exception) {
	// retrieve existing user
	existingUserDetail, exc := user.Get("code", code)
	if exc != nil {
		return nil, exc
	}
	// update user data
	DB.Model(&existingUserDetail).Updates(data)

	return existingUserDetail, nil
}

// ------------------------------------------------------

/**
 * Delete Existing User Record
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
func (user *User) Delete(code string) (string, exception.Exception) {
	// check if code exists
	var existingUser User
	query := DB.Where(
		map[string]interface{}{
			"code": code,
		},
	).First(&existingUser)
	if query.RecordNotFound() {
		return "", exception.NewRecordNotFound("User")
	}

	// get user code
	userCode := existingUser.Code

	// delete user
	DB.Delete(existingUser)

	return userCode, nil
}

/**
 * Get User Detail by Email
 *
 * @param 	string email
 *
 * @return 	*User
 * @return 	exception.Exception
 */
func (user *User) GetByEmail(email string) (*User, exception.Exception) {
	// retrieve user detail
	var userDetail User
	query := DB.Where(
		map[string]interface{}{
			"email": email,
		},
	).First(&userDetail)

	// if record not found
	if query.RecordNotFound() {
		return nil, exception.NewRecordNotFound("")
	}

	return &userDetail, nil
}

// ------------------------------------------------------

/**
 * Get User Detail by Token
 *
 * @param 	string email
 *
 * @return 	*User
 * @return 	exception.Exception
 */
func (user *User) GetByToken(token string) (*User, exception.Exception) {
	// retrieve user detail
	var userDetail User
	query := DB.Where(
		map[string]interface{}{
			"api_token": token,
			"status":    "ACTIVE",
		},
	).Where("api_token_expiration > ?", time.Now()).First(&userDetail)

	// if record not found
	if query.RecordNotFound() {
		return nil, exception.NewRecordNotFound("")
	}

	return &userDetail, nil
}

// ------------------------------------------------------

/**
 * Update User Token
 */
func (user *User) UpdateToken() {
	// update api token and the expiration
	expirationMinutes, _ := strconv.Atoi(os.Getenv("API_TOKEN_EXPIRATION_DURATION"))
	DB.Model(&user).Updates(
		map[string]interface{}{
			"api_token":            hash.Randomize(),
			"api_token_expiration": time.Now().Add(time.Minute * time.Duration(expirationMinutes)),
		},
	)
}
