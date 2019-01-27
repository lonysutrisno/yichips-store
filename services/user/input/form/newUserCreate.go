package data

import (
	// native
	"strings"

	// yichips
	"yichips/utils/helper/hash"
	"yichips/utils/helper/random"
)

type NewUserCreate struct {
	Name     string
	Code     string
	Email    string
	Password string
	Phone    string
	Role     string
	Status   string
}

func NewNewUserCreate(data map[string]interface{}) NewUserCreate {
	return NewUserCreate{
		Name:  data["Name"].(string),
		Email: data["Email"].(string),
		Phone: data["Phone"].(string),
		Role:  "USER",

		Password: hash.Make(data["Password"].(string)),
		Code:     "ADM-" + strings.ToUpper(random.String(7)),
		Status:   "VERIFY",
	}
}
