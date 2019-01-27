package output

import (
	// yichips
	"yichips/utils/collection"
	"yichips/utils/contract"

	// product service
	"yichips/services/product/model"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type ProductDetail struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Qty         int     `json:"qty"`
	Price       float64 `json:"price"`
	Weight      int     `json:"weight"`
	Images      string  `json:"images"`
	Description string  `json:"description"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Product Detail Output Model
 *
 * @param 	*model.Product 	product
 *
 * @return 	contract.ModelResponse
 */
func NewProductDetail(product *model.Product) contract.ModelResponse {
	return ProductDetail{
		Code:        product.Code,
		Name:        product.Name,
		Qty:         product.Qty,
		Price:       product.Price,
		Weight:      product.Weight,
		Images:      product.Images,
		Description: product.Description,
	}
}

// ------------------------------------------------------

/**
 * Product List Output Models
 *
 * @param 	[]model.Product 	product
 *
 * @return 	collection.ModelResponse
 */
func NewProductDetailList(productList []model.Product) collection.ModelResponse {
	productDetails := collection.NewModelResponse()
	for _, product := range productList {
		productDetails.Add(NewProductDetail(&product))
	}

	return productDetails
}
