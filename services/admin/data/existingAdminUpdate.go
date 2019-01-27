package data

import (
	// yichips
	"yichips/utils/helper/hash"
)

type ExistingAdminUpdate struct {
	Name     string
	Email    string
	Password string
	Status   string
}

func NewExistingAdminUpdate(data map[string]interface{}) ExistingAdminUpdate {
	return ExistingAdminUpdate{
		Name:   data["Name"].(string),
		Email:  data["Email"].(string),
		Status: data["Status"].(string),

		Password: hash.Make(data["Password"].(string)),
	}
}
