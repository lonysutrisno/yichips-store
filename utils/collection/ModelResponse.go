package collection

import (
	// yichips
	"yichips/utils/contract"
	"yichips/utils/rest/meta"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type ModelResponse struct {
	List []contract.ModelResponse `json:"list"`
	Meta meta.Meta                `json:"meta"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Constructor
 *
 * @return 	ModelResponse
 */
func NewModelResponse() ModelResponse {
	return ModelResponse{List: make([]contract.ModelResponse, 0)}
}

// ------------------------------------------------------

/**
 * Add New Item
 *
 * @param 	contract.ModelResponse 	newModel
 */
func (model *ModelResponse) Add(newModel contract.ModelResponse) {
	model.List = append(model.List, newModel)
}

// ------------------------------------------------------

/**
 * Get List Length
 *
 * @return int
 */
func (model *ModelResponse) Length() int {
	return len(model.List)
}

/**
 * Add List Meta
 *
 * @param 	int 	currentPage
 * @param 	int 	listLimit
 *
 * @return 	*ModelResponse
 */
func (model *ModelResponse) WithMeta(currentPage int, listLimit int, listTotal int) *ModelResponse {
	model.Meta = meta.NewListMeta(currentPage, listLimit, model.Length(), listTotal)

	return model
}
