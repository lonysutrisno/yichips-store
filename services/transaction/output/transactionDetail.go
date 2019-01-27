package output

import (
	"fmt"
	// yichips
	"yichips/utils/collection"
	"yichips/utils/contract"

	// transaction service
	"yichips/services/transaction/model"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Transaction struct {
	TransactionCode string  `json:"transaction_code"`
	TotalPrice      float64 `json:"total_price"`
	Address         string  `json:"address"`
	Phone           string  `json:"phone"`
	DeliveryType    string  `json:"delivery_type"`
	NoResi          string  `json:"no_resi"`
}
type TransactionDetail struct {
	TransactionID int     `json:"transaction_id"`
	ProductID     int     `json:"product_id"`
	Qty           int     `json:"qty"`
	Price         float64 `json:"price"`
	Transaction   Transaction
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Transaction Detail Output Model
 *
 * @param 	*model.Transaction 	transaction
 *
 * @return 	contract.ModelResponse
 */
func NewTransactionDetail(transaction *model.Transaction) contract.ModelResponse {
	return Transaction{
		TransactionCode: transaction.TransactionCode,
		TotalPrice:      transaction.TotalPrice,
		Address:         transaction.Address,
		Phone:           transaction.Phone,
		DeliveryType:    transaction.DeliveryType,
		NoResi:          transaction.NoResi,
	}
}

// ------------------------------------------------------

/**
 * Transaction List Output Models
 *
 * @param 	[]model.Transaction 	transaction
 *
 * @return 	collection.ModelResponse
 */
func NewTransactionDetailList(transactionList []model.Transaction) collection.ModelResponse {
	transactionDetails := collection.NewModelResponse()
	fmt.Println()
	for _, transaction := range transactionList {
		transactionDetails.Add(NewTransactionDetail(&transaction))
	}

	return transactionDetails
}
