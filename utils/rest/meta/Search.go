package meta

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Search struct {
	Searchables []string `json:"searchables"`
	Active      string   `json:"active"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

func NewMetaSearch(searchables []string, active string) Search {
	return Search{
		Searchables: searchables,
		Active:      active,
	}
}
