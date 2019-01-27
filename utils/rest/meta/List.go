package meta

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type List struct {
	Current int `json:"current"`
	Limit   int `json:"limit"`
	Total   int `json:"total"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

func NewMetaList(current int, limit int, total int) List {
	return List{
		Current: current,
		Limit:   limit,
		Total:   total,
	}
}
