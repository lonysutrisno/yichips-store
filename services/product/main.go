package product

import (

	// native
	// "strings"

	// yichips

	"strconv"
	"yichips/config"
	"yichips/exception"
	"yichips/utils/contract"
	"yichips/utils/rest"

	// product service
	input "yichips/services/product/input/form"
	request "yichips/services/product/input/requestValidator"
	"yichips/services/product/model"
	"yichips/services/product/output"

	"github.com/labstack/echo"
)

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Product Service
 |	------------------
 |
 |	This service is used to handle product management operations
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------

var ProductModel = model.NewProduct()

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Main struct {
}

var productListLimit, _ = strconv.Atoi(config.App["admin_list_limit"])

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Handle Product Create
 */
func (main *Main) Create(e echo.Context) (err error) {

	// get request and validate
	req := new(request.Create)
	e.Bind(req)
	if err = e.Validate(req); err != nil {
		return rest.ConstructErrorResponse(e, exception.NewInputValidationFailed(err.Error()))
	}
	// map req to input data
	reqData := input.NewNewProductCreate(
		map[string]interface{}{
			"Name":        req.Name,
			"Qty":         req.Qty,
			"Price":       req.Price,
			"Weight":      req.Weight,
			"Images":      req.Images,
			"Description": req.Description,
		},
	)
	//insert data to db
	product, exc := ProductModel.Create(reqData)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}
	data := map[string]contract.Model{
		"created_product": product,
	}
	return rest.ConstructSuccessResponse(e, data)
}

/**
 * Handle Product List
 */
func (main *Main) List(e echo.Context) (err error) {
	// get page
	page, _ := strconv.Atoi(e.QueryParam("page"))
	// get product list
	products, exc := ProductModel.Fetch(((page - 1) * productListLimit), productListLimit)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// get total list
	total, exc := ProductModel.Count()
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	productList := output.NewProductDetailList(products)

	data := map[string]contract.Model{
		"products": productList.WithMeta(page, productListLimit, total),
	}

	return rest.ConstructSuccessResponse(e, data)
}

/**
 * Handle Product Detail Show
 */
func (main *Main) GetDetail(e echo.Context) (err error) {
	// get path parameter
	productCode := e.Param("code")

	// get product details
	productDetail, exc := ProductModel.Get("code", productCode)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	data := map[string]contract.Model{
		"product_detail": output.NewProductDetail(productDetail),
	}

	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------

/**
 * Create New Product
 *
 * @param 	dataRaw.NewProductCreate 	data
 *
 * @return 	contract.ServiceOutput
 * @return 	exception.Exception
 */
// func (main *Main) CreateNew(data dataRaw.NewProductCreate) (contract.ServiceOutput, exception.Exception) {
// 	// create product
// 	product, exc := ProductModel.Create(data)
// 	if exc != nil {
// 		return nil, exc
// 	}

// 	return output.NewProductDetail(product), nil
// }

// ------------------------------------------------------

/**
 * Delete Existing Product
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
// func (main *Main) DeleteExisting(code string) (string, exception.Exception) {
// 	// delete product data by code
// 	productCode, exc := ProductModel.Delete(code)
// 	if exc != nil {
// 		return "", exc
// 	}

// 	return productCode, nil
// }

// ------------------------------------------------------
