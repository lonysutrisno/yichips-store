package requestValidator

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Create struct {
	UserID    int `json:"user_id" validate:"required"`
	ProductID int `json:"product_id" validate:"required"`
	Qty       int `json:"qty" validate:"required"`
}
