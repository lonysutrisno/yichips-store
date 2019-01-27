package collection

import (
	// yichips
	"yichips/utils/contract"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Model struct {
	List []contract.Model `json:"list"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Constructor
 *
 * @return 	Model
 */
func NewModel() Model {
	return Model{List: make([]contract.Model, 0)}
}

// ------------------------------------------------------

/**
 * Add New Item
 *
 * @param 	contract.Model 	newModel
 */
func (model *Model) Add(newModel contract.Model) {
	model.List = append(model.List, newModel)
}

// ------------------------------------------------------

/**
 * Get List Length
 *
 * @return 	int
 */
func (model *Model) Length() int {
	return len(model.List)
}
