package model

import (
	// external

	"github.com/jinzhu/copier"

	// yichips
	"yichips/bootstrap"
	base "yichips/core/model"
	"yichips/exception"

	// product service
	productData "yichips/services/product/input/form"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Product struct {
	base.Product `mapstructure:",squash"`
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
 * @return 	Product
 */
func NewProduct() Product {
	return Product{}
}

// ------------------------------------------------------

/**
 * Fetch Product Records
 *
 * @return 	[]Product
 * @return 	exception.Exception
 */
func (product *Product) Fetch(offset int, limit int) ([]Product, exception.Exception) {
	var productList []Product
	DB.Offset(offset).Limit(limit).Find(&productList)

	return productList, nil
}

// ------------------------------------------------------

/**
 * Count Total Product Records
 *
 * @return 	int
 * @return 	exception.Exception
 */
func (product *Product) Count() (int, exception.Exception) {
	var productList []Product
	var total int
	DB.Find(&productList).Count(&total)

	return total, nil
}

// ------------------------------------------------------

/**
 * Get Product Detail by Code
 *
 * @param 	string code
 *
 * @return 	*Product
 * @return 	exception.Exception
 */
func (product *Product) Get(field, code string) (*Product, exception.Exception) {
	// retrieve product detail
	var productDetail Product
	query := DB.Where(
		map[string]interface{}{
			field: code,
		},
	).First(&productDetail)

	// if record not found
	if query.RecordNotFound() {
		return nil, exception.NewRecordNotFound("Product")
	}

	return &productDetail, nil
}

// ------------------------------------------------------

/**
 * Store New Product Record
 *
 * @param 	productData.NewProductCreate 	data
 *
 * @return 	*Product
 * @return 	exception.Exception
 */
func (product *Product) Create(data productData.NewProductCreate) (*Product, exception.Exception) {
	// check if name already has been used
	var existingProduct Product
	query := DB.Where(
		map[string]interface{}{
			"name": data.Name,
		},
	).First(&existingProduct)
	if !query.RecordNotFound() {
		return nil, exception.NewDuplicateRecordFound("Name")
	}

	// store product data
	var newProduct Product
	copier.Copy(&newProduct, &data)
	DB.Create(&newProduct)

	return &newProduct, nil
}

// ------------------------------------------------------

/**
 * Update Existing Product Record
 *
 * @param 	string 	code
 * @param 	map[string]interface{} 	data
 *
 * @return 	*Product
 * @return 	exception.Exception
 */
func (product *Product) Update(code string, data map[string]interface{}) (*Product, exception.Exception) {
	// retrieve existing product
	existingProductDetail, exc := product.Get("code", code)
	if exc != nil {
		return nil, exc
	}
	// update product data
	DB.Model(&existingProductDetail).Updates(data)

	return existingProductDetail, nil
}

// ------------------------------------------------------

/**
 * Delete Existing Product Record
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
func (product *Product) Delete(code string) (string, exception.Exception) {
	// check if code exists
	var existingProduct Product
	query := DB.Where(
		map[string]interface{}{
			"code": code,
		},
	).First(&existingProduct)
	if query.RecordNotFound() {
		return "", exception.NewRecordNotFound("Product")
	}

	// get product code
	productCode := existingProduct.Code

	// delete product
	DB.Delete(existingProduct)

	return productCode, nil
}
