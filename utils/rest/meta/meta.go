package meta

type Meta struct {
	Page   Page   `json:"page"`
	List   List   `json:"list"`
	Filter Filter `json:"filter"`
	Search Search `json:"search"`
	Sort   Sort   `json:"sort"`
}

func NewListMeta(currentPage int, listLimit int, currentList int, totalList int) Meta {
	// get total page
	partialPage := 1
	if (totalList % listLimit) == 0 {
		partialPage = 0
	}
	totalPage := (totalList / listLimit) + partialPage

	// adjust inputted current list
	if currentPage < totalPage {
		currentList = listLimit
	}

	// adjust inputted current page
	if currentPage > totalPage {
		currentPage = totalPage
	} else if currentPage < 1 {
		currentPage = 1
	}

	return Meta{
		Page: NewMetaPage(currentPage, totalPage),
		List: NewMetaList(currentList, listLimit, totalList),
		// [TO DO]
		Filter: NewMetaFilter(
			map[string][]string{
				"name": {
					"abc",
					"def",
					"ghi",
				},
				"code": {
					"123",
					"456",
				},
			},
			map[string]string{
				"name": "abc;def",
			},
		),
		// [TO DO]
		Search: NewMetaSearch(
			[]string{
				"abc",
				"def",
				"ghi",
			},
			"qwerty",
		),
		// [TO DO]
		Sort: NewMetaSort(
			[]string{
				"name",
				"code",
			},
			"-name",
		),
	}
}
