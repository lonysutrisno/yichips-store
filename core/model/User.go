package model

import (
	"github.com/jinzhu/gorm"

	"time"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type User struct {
	gorm.Model
	Carts              []Cart
	Code               string `gorm:"type:varchar(20);not null"`
	Name               string `gorm:"type:varchar(150);not null"`
	Email              string `gorm:"type:varchar(70);not null"`
	Phone              string `gorm:"type:varchar(70);not null"`
	Password           string `gorm:"type:varchar(255);not null"`
	Address            string `gorm:"type:varchar(255);default:null"`
	PostalCode         string `gorm:"type:varchar(255);default:null"`
	FbToken            string `gorm:"type:varchar(255);default:null"`
	GoogleToken        string `gorm:"type:varchar(255);default:null"`
	Status             string `gorm:"type:varchar(20);default:'ACTIVE';not null"`
	Role               string `gorm:"type:varchar(255);not null"`
	ApiToken           string `gorm:"type:varchar(255);not null"`
	ApiTokenExpiration *time.Time
}
