package data

import (
	// native
	"strings"

	// yichips
	"yichips/utils/helper/hash"
	"yichips/utils/helper/random"
)

type NewAdminCreate struct {
	Name     string
	Code     string
	Email    string
	Password string
	Status   string
}

func NewNewAdminCreate(data map[string]interface{}) NewAdminCreate {
	return NewAdminCreate{
		Name:  data["Name"].(string),
		Email: data["Email"].(string),

		Password: hash.Make(data["Password"].(string)),
		Code:     "USR-" + strings.ToUpper(random.String(7)),
		Status:   "ACTIVE",
	}
}
