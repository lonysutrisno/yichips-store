package model

import (
	"github.com/jinzhu/gorm"

	"time"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Admin struct {
	gorm.Model
	Code               string `gorm:"type:varchar(20);not null"`
	Name               string `gorm:"type:varchar(150);not null"`
	Email              string `gorm:"type:varchar(70);not null"`
	Password           string `gorm:"type:varchar(255);not null"`
	Status             string `gorm:"type:varchar(20);default:'ACTIVE';not null"`
	ApiToken           string `gorm:"type:varchar(255);not null"`
	ApiTokenExpiration *time.Time
}
