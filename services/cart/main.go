package cart

import (

	// native
	// "strings"

	// yichips

	"strconv"
	"yichips/bootstrap"
	"yichips/config"
	"yichips/exception"
	"yichips/utils/contract"
	"yichips/utils/rest"

	// cart service
	base "yichips/core/model"
	request "yichips/services/cart/input/requestValidator"
	Cart "yichips/services/cart/model"
	Product "yichips/services/product/model"
	User "yichips/services/user/model"

	"github.com/labstack/echo"
)

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Cart Service
 |	------------------
 |
 |	This service is used to handle cart management operations
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------

var CartModel = Cart.NewCart()
var ProductModel = Product.NewProduct()
var UserModel = User.NewUser()

type Carts struct {
	base.Cart `mapstructure:",squash"`
}
type Users struct {
	base.User `mapstructure:",squash"`
}

var DB = bootstrap.DB

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Main struct {
}

var cartListLimit, _ = strconv.Atoi(config.App["admin_list_limit"])

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Handle Cart Create
 */
func (main *Main) Create(e echo.Context) (err error) {

	// get request and validate
	req := new(request.Create)
	e.Bind(req)
	if err = e.Validate(req); err != nil {
		return rest.ConstructErrorResponse(e, exception.NewInputValidationFailed(err.Error()))
	}

	// get user details
	userId := strconv.Itoa(req.UserID)
	_, excUser := UserModel.Get("id", userId)
	if excUser != nil {
		return rest.ConstructErrorResponse(e, excUser)
	}

	// get product details
	productId := strconv.Itoa(req.ProductID)
	_, excProduct := ProductModel.Get("id", productId)
	if excProduct != nil {
		return rest.ConstructErrorResponse(e, excProduct)
	}
	//check if product exist in cart

	if CartModel.CheckExistingProduct(userId, productId) == true {
		rest.ConstructErrorResponse(e, exception.NewDuplicateRecordFound("product"))
	} else {
		// map req to input data
		reqData :=
			map[string]interface{}{
				"UserID":    req.UserID,
				"ProductID": req.ProductID,
				"Qty":       req.Qty,
			}

		//insert data to db
		cart, exc := CartModel.Create(reqData)
		if exc != nil {
			return rest.ConstructErrorResponse(e, exc)
		}
		data := map[string]contract.Model{
			"created_cart": cart,
		}
		return rest.ConstructSuccessResponse(e, data)
	}
	return
}

/**
 * Handle Cart Detail Show
 */
func (main *Main) GetCartByUserId(e echo.Context) (err error) {
	UserID := e.Param("code")

	user, cart, exc := CartModel.GetAll("user_id", UserID)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	data := map[string]contract.Model{
		"cart_detail": cart,
		"user_detail": user,
	}

	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------

/**
 * Update Item Cart
 *
 * @param 	dataRaw.NewCartCreate 	data
 *
 * @return 	contract.ServiceOutput
 * @return 	exception.Exception
 */
func (main *Main) UpdateCart(e echo.Context) (err error) {
	id := e.Param("code")
	req := new(request.Update)
	e.Bind(req)
	qty, _ := strconv.Atoi(req.Qty)
	if err = e.Validate(req); err != nil {
		return rest.ConstructErrorResponse(e, exception.NewInputValidationFailed(err.Error()))
	}
	// update cart
	user, cart, exc := CartModel.Update(id, map[string]interface{}{
		"Qty": qty,
	})
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}
	// prepare data
	data := map[string]contract.Model{
		"cart_detail": cart,
		"user_detail": user,
	}

	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------

/**
 * Delete Existing Cart
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
func (main *Main) RemoveProduct(e echo.Context) (err error) {
	id := e.Param("code")
	// delete cart data by code
	user, cart, exc := CartModel.Delete(id)

	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}
	// prepare data
	data := map[string]contract.Model{
		"cart_detail": cart,
		"user_detail": user,
	}
	return rest.ConstructSuccessResponse(e, data)
}

/**
 * Delete Existing Cart
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
func (main *Main) EmptyCart(e echo.Context) (err error) {
	id := e.Param("code")
	// delete cart data by code
	user, cart, exc := CartModel.DeleteAll(id)

	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}
	// prepare data
	data := map[string]contract.Model{
		"cart_detail": cart,
		"user_detail": user,
	}
	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------
