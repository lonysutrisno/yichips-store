package model

import (
	"github.com/jinzhu/gorm"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Cart struct {
	gorm.Model
	Product   Product
	UserID    int `gorm:"type:int;not null"`
	ProductID int `gorm:"type:int;not null"`
	Qty       int `gorm:"type:int;not null"`
}
