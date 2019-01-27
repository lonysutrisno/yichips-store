package data

import (
	// native
	"strings"

	// yichips

	"yichips/utils/helper/random"
)

type NewProductCreate struct {
	Code        string
	Name        string
	Qty         int
	Price       float64
	Weight      int
	Images      string
	Description string
}

func NewNewProductCreate(data map[string]interface{}) NewProductCreate {
	return NewProductCreate{
		Name:        data["Name"].(string),
		Qty:         data["Qty"].(int),
		Price:       data["Price"].(float64),
		Weight:      data["Weight"].(int),
		Description: data["Description"].(string),
		Images:      data["Images"].(string),

		Code: "YCP-" + strings.ToUpper(random.String(7)),
	}
}
