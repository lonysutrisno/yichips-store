package meta

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Filter struct {
	Filterables map[string][]string `json:"filterables"`
	Active      map[string]string   `json:"active"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

func NewMetaFilter(filterables map[string][]string, active map[string]string) Filter {
	return Filter{
		Filterables: filterables,
		Active:      active,
	}
}
