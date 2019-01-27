package requestValidator

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Create struct {
	Name        string  `json:"name" validate:"required"`
	Qty         int     `json:"qty" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Weight      int     `json:"weight" validate:"required"`
	Images      string  `json:"images" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
