package meta

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Page struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

func NewMetaPage(current int, total int) Page {
	return Page{
		Current: current,
		Total:   total,
	}
}
