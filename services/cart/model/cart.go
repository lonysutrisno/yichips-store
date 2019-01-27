package model

import (
	"fmt"
	"strconv"

	// external

	"github.com/jinzhu/copier"

	// yichips
	"yichips/bootstrap"
	base "yichips/core/model"
	"yichips/exception"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Cart struct {
	base.Cart `mapstructure:",squash"`
}
type Product struct {
	base.Product `mapstructure:",squash"`
}
type User struct {
	base.User `mapstructure:",squash"`
}
type New struct {
	UserID    int
	ProductID int
	Qty       int
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
 * @return 	Cart
 */
func NewCart() Cart {
	return Cart{}
}

// ------------------------------------------------------

/**
 * Fetch Cart Records
 *
 * @return 	[]Cart
 * @return 	exception.Exception
 */
func (cart *Cart) Fetch(offset int, limit int) ([]Cart, exception.Exception) {
	var cartList []Cart
	DB.Offset(offset).Limit(limit).Find(&cartList)

	return cartList, nil
}

// ------------------------------------------------------

/**
 * Count Total Cart Records
 *
 * @return 	int
 * @return 	exception.Exception
 */
func (cart *Cart) Count() (int, exception.Exception) {
	var cartList []Cart
	var total int
	DB.Find(&cartList).Count(&total)

	return total, nil
}

// ------------------------------------------------------

/**
 * Get Cart Detail by Code
 *
 * @param 	string code
 *
 * @return 	*Cart
 * @return 	exception.Exception
 */
func (cart *Cart) Get(field, code string) (*Cart, exception.Exception) {
	// retrieve cart detail
	var cartDetail Cart
	query := DB.Where(
		map[string]interface{}{
			field: code,
		},
	).First(&cartDetail)

	// if record not found
	if query.RecordNotFound() {
		return nil, exception.NewRecordNotFound("Cart")
	}

	return &cartDetail, nil
}

/**
 * Check Existing Product
 *
 * @param 	string code
 *
 * @return 	*Cart
 * @return 	exception.Exception
 */
func (cart *Cart) CheckExistingProduct(userId, productId string) bool {
	// retrieve cart detail
	var cartDetail Cart
	query := DB.Where(
		map[string]interface{}{
			"user_id":    userId,
			"product_id": productId,
		},
	).First(&cartDetail)

	// if record not found
	if query.RecordNotFound() {
		return false
	}

	return true
}

/**
 * Get Cart Detail by Code
 *
 * @param 	string code
 *
 * @return 	*Cart
 * @return 	exception.Exception
 */
func (cart *Cart) GetAll(field, code string) (*User, []Cart, exception.Exception) {
	// retrieve cart detail
	var cartDetail []Cart
	// var products []Product
	var userDetail User

	queryUser := DB.Where(
		map[string]interface{}{
			"id": code,
		},
	).Find(&userDetail)
	if queryUser.RecordNotFound() {
		return nil, nil, exception.NewRecordNotFound("User")
	}

	queryProduct := DB.Where(
		map[string]interface{}{
			field: code,
		},
	).Preload("Product").Find(&cartDetail)
	if queryProduct.RecordNotFound() {
		return nil, nil, exception.NewRecordNotFound("Product")
	}

	return &userDetail, cartDetail, nil
}

// ------------------------------------------------------

/**
 * Store New Cart Record
 *
 * @param 	cartData.NewCartCreate 	data
 *
 * @return 	*Cart
 * @return 	exception.Exception
 */
func (cart *Cart) Create(data map[string]interface{}) (*Cart, exception.Exception) {

	// store cart data
	var newCart Cart
	var x = New{
		UserID:    data["UserID"].(int),
		ProductID: data["ProductID"].(int),
		Qty:       data["Qty"].(int),
	}

	copier.Copy(&newCart, &x)
	DB.Create(&newCart)

	return &newCart, nil
}

// ------------------------------------------------------

/**
 * Update Existing Cart Record
 *
 * @param 	string 	code
 * @param 	map[string]interface{} 	data
 *
 * @return 	*Cart
 * @return 	exception.Exception
 */
func (cart *Cart) Update(code string, data map[string]interface{}) (*User, []Cart, exception.Exception) {
	// retrieve existing cart

	var cartDetail []Cart

	existingCartDetail, exc := cart.Get("id", code)
	if exc != nil {
		return nil, nil, exc
	}
	// update cart data
	DB.Model(&existingCartDetail).Updates(data)
	userDetail, cartDetail, excCart := cart.GetAll("user_id", strconv.Itoa(existingCartDetail.UserID))
	if exc != nil {
		return nil, nil, excCart
	}
	return userDetail, cartDetail, nil

}

// ------------------------------------------------------

/**
 * Delete Existing Cart Record
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
func (cart *Cart) Delete(code string) (*User, []Cart, exception.Exception) {
	// check if code exists
	var existingCart Cart
	query := DB.Where(
		map[string]interface{}{
			"ID": code,
		},
	).First(&existingCart)

	if query.RecordNotFound() {
		return nil, nil, exception.NewRecordNotFound("Cart")
	}

	UserID := existingCart.UserID

	DB.Unscoped().Delete(&existingCart)

	userDetail, cartDetail, excCart := cart.GetAll("user_id", strconv.Itoa(UserID))

	if excCart != nil {
		return nil, nil, excCart
	}

	return userDetail, cartDetail, nil
}

/**
 * Delete Existing Cart Record
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
func (cart *Cart) DeleteAll(code string) (*User, []Cart, exception.Exception) {
	// check if code exists
	var existingCart []Cart
	query := DB.Where(map[string]interface{}{
		"user_id": code,
	}).First(&existingCart)

	if query.RecordNotFound() {
		return nil, nil, exception.NewRecordNotFound("Cart")
	}
	fmt.Println(existingCart)
	DB.Unscoped().Delete(existingCart)

	userDetail, cartDetail, excCart := cart.GetAll("user_id", code)

	if excCart != nil {
		return nil, nil, excCart
	}

	return userDetail, cartDetail, nil
}
