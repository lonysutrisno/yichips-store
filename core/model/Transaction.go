package model

import (
	"github.com/jinzhu/gorm"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Transaction struct {
	gorm.Model
	TransactionCode   string  `gorm:"type:varchar(150);not null"`
	TotalPrice        float64 `gorm:"type:float;not null"`
	Address           string  `gorm:"type:varchar(150);not null"`
	Phone             string  `gorm:"type:varchar(150);not null"`
	DeliveryType      string  `gorm:"type:varchar(150);not null"`
	NoResi            string  `gorm:"type:varchar(150);not null"`
	TransactionDetail []TransactionDetail
}
