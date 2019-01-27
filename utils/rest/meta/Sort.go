package meta

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Sort struct {
	Sortables []string `json:"sortables"`
	Active    string   `json:"active"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

func NewMetaSort(sortables []string, active string) Sort {
	return Sort{
		Sortables: sortables,
		Active:    active,
	}
}
