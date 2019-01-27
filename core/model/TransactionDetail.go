package model

import (
	"github.com/jinzhu/gorm"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type TransactionDetail struct {
	gorm.Model
	TransactionID int     `gorm:"type:int;not null"`
	ProductID     int     `gorm:"type:int;not null"`
	Qty           int     `gorm:"type:int;not null"`
	Price         float64 `gorm:"type:float;not null"`
	Transaction   Transaction
}
