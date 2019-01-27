package model

import (
	// external

	"github.com/jinzhu/copier"

	// yichips
	"yichips/bootstrap"
	base "yichips/core/model"
	"yichips/exception"

	// transaction service
	transactionData "yichips/services/transaction/input/form"
)

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Transaction struct {
	base.Transaction `mapstructure:",squash"`
}

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------

var DB = bootstrap.DB

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Constructor
 *
 * @return 	Transaction
 */
func NewTransaction() Transaction {
	return Transaction{}
}

// ------------------------------------------------------

/**
 * Fetch Transaction Records
 *
 * @return 	[]Transaction
 * @return 	exception.Exception
 */
func (transaction *Transaction) Fetch(offset int, limit int) ([]Transaction, exception.Exception) {
	var transactionList []Transaction
	DB.Offset(offset).Limit(limit).Find(&transactionList)

	return transactionList, nil
}

// ------------------------------------------------------

/**
 * Count Total Transaction Records
 *
 * @return 	int
 * @return 	exception.Exception
 */
func (transaction *Transaction) Count() (int, exception.Exception) {
	var transactionList []Transaction
	var total int
	DB.Find(&transactionList).Count(&total)

	return total, nil
}

// ------------------------------------------------------

/**
 * Get Transaction Detail by Code
 *
 * @param 	string code
 *
 * @return 	*Transaction
 * @return 	exception.Exception
 */
func (transaction *Transaction) Get(field, code string) (*Transaction, exception.Exception) {
	// retrieve transaction detail
	var transactionDetail Transaction
	query := DB.Where(
		map[string]interface{}{
			field: code,
		},
	).First(&transactionDetail)

	// if record not found
	if query.RecordNotFound() {
		return nil, exception.NewRecordNotFound("Transaction")
	}

	return &transactionDetail, nil
}

// ------------------------------------------------------

/**
 * Store New Transaction Record
 *
 * @param 	transactionData.NewTransactionCreate 	data
 *
 * @return 	*Transaction
 * @return 	exception.Exception
 */
func (transaction *Transaction) Create(data transactionData.NewTransactionCreate) (*Transaction, exception.Exception) {
	// check if name already has been used
	// var existingTransaction Transaction
	// query := DB.Where(
	// 	map[string]interface{}{
	// 		"name": data.Name,
	// 	},
	// ).First(&existingTransaction)
	// if !query.RecordNotFound() {
	// 	return nil, exception.NewDuplicateRecordFound("Name")
	// }

	// store transaction data
	var newTransaction Transaction
	copier.Copy(&newTransaction, &data)
	DB.Create(&newTransaction)

	return &newTransaction, nil
}

// ------------------------------------------------------

/**
 * Update Existing Transaction Record
 *
 * @param 	string 	code
 * @param 	map[string]interface{} 	data
 *
 * @return 	*Transaction
 * @return 	exception.Exception
 */
// func (transaction *Transaction) Update(code string, data map[string]interface{}) (*Transaction, exception.Exception) {
// 	// retrieve existing transaction
// 	existingTransactionDetail, exc := transaction.Get("code", code)
// 	if exc != nil {
// 		return nil, exc
// 	}
// 	// update transaction data
// 	DB.Model(&existingTransactionDetail).Updates(data)

// 	return existingTransactionDetail, nil
// }

// ------------------------------------------------------

/**
 * Delete Existing Transaction Record
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
// func (transaction *Transaction) Delete(code string) (string, exception.Exception) {
// 	// check if code exists
// 	var existingTransaction Transaction
// 	query := DB.Where(
// 		map[string]interface{}{
// 			"code": code,
// 		},
// 	).First(&existingTransaction)
// 	if query.RecordNotFound() {
// 		return "", exception.NewRecordNotFound("Transaction")
// 	}

// 	// get transaction code
// 	transactionCode := existingTransaction.Code

// 	// delete transaction
// 	DB.Delete(existingTransaction)

// 	return transactionCode, nil
// }
