package data

import (
	// native
	"strings"

	// yichips

	"yichips/utils/helper/random"
)

type NewTransactionCreate struct {
	TransactionCode string
	Qty             int
	TotalPrice      float64
	Address         string
	Phone           string
	DeliveryType    string
	NoResi          string
}

type NewTransactionDetailCreate struct {
	TransactionCode string
	Qty             int
	TotalPrice      float64
	Address         string
	Phone           string
	DeliveryType    string
	NoResi          string
}

func NewNewTransactionCreate(data map[string]interface{}) NewTransactionCreate {
	return NewTransactionCreate{
		Qty:             data["Qty"].(int),
		TotalPrice:      data["TotalPrice"].(float64),
		Address:         data["Address"].(string),
		DeliveryType:    data["DeliveryType"].(string),
		NoResi:          data["NoResi"].(string),
		TransactionCode: "TRX-" + strings.ToUpper(random.String(7)),
	}
}
