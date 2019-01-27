package requestValidator

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Update struct {
	Qty string `json:"qty" validate:"required"`
}
