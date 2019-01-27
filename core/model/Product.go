package model

import (
	"github.com/jinzhu/gorm"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Product struct {
	gorm.Model
	// Carts       []Cart
	Name        string  `gorm:"type:varchar(150);not null"`
	Code        string  `gorm:"type:varchar(20);not null"`
	Price       float64 `gorm:"type:float;not null"`
	Qty         int     `gorm:"type:int;not null"`
	Description string  `gorm:"type:varchar(350);not null"`
	Weight      int     `gorm:"type:int;not null"`
	Images      string  `sql:"type:text;"`
}
