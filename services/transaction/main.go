package transaction

import (

	// native
	// "strings"

	// yichips

	"strconv"
	"yichips/config"
	"yichips/exception"
	"yichips/utils/contract"
	"yichips/utils/rest"

	// transaction servigo get ce
	input "yichips/services/transaction/input/form"
	request "yichips/services/transaction/input/requestValidator"
	"yichips/services/transaction/model"
	"yichips/services/transaction/output"

	"github.com/labstack/echo"
)

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Transaction Service
 |	------------------
 |
 |	This service is used to handle transaction management operations
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// DECLARATIONS
// ------------------------------------------------------

var TransactionModel = model.NewTransaction()

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN STRUCT
// ------------------------------------------------------

type Main struct {
}

var transactionListLimit, _ = strconv.Atoi(config.App["admin_list_limit"])

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Handle Transaction Create
 */
func (main *Main) Create(e echo.Context) (err error) {

	// get request and validate
	req := new(request.Create)
	e.Bind(req)
	if err = e.Validate(req); err != nil {
		return rest.ConstructErrorResponse(e, exception.NewInputValidationFailed(err.Error()))
	}
	// map req to input data
	reqData := input.NewNewTransactionCreate(
		map[string]interface{}{
			"Name":        req.Name,
			"Qty":         req.Qty,
			"Price":       req.Price,
			"Weight":      req.Weight,
			"Images":      req.Images,
			"Description": req.Description,
		},
	)
	//insert data to db
	transaction, exc := TransactionModel.Create(reqData)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}
	data := map[string]contract.Model{
		"created_transaction": transaction,
	}
	return rest.ConstructSuccessResponse(e, data)
}

/**
 * Handle Transaction List
 */
func (main *Main) List(e echo.Context) (err error) {
	// get page
	page, _ := strconv.Atoi(e.QueryParam("page"))
	// get transaction list
	transactions, exc := TransactionModel.Fetch(((page - 1) * transactionListLimit), transactionListLimit)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// get total list
	total, exc := TransactionModel.Count()
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	transactionList := output.NewTransactionDetailList(transactions)

	data := map[string]contract.Model{
		"transactions": transactionList.WithMeta(page, transactionListLimit, total),
	}

	return rest.ConstructSuccessResponse(e, data)
}

/**
 * Handle Transaction Detail Show
 */
func (main *Main) GetDetail(e echo.Context) (err error) {
	// get path parameter
	transactionCode := e.Param("code")

	// get transaction details
	transactionDetail, exc := TransactionModel.Get("code", transactionCode)
	if exc != nil {
		return rest.ConstructErrorResponse(e, exc)
	}

	// prepare data
	data := map[string]contract.Model{
		"transaction_detail": output.NewTransactionDetail(transactionDetail),
	}

	return rest.ConstructSuccessResponse(e, data)
}

// ------------------------------------------------------

/**
 * Create New Transaction
 *
 * @param 	dataRaw.NewTransactionCreate 	data
 *
 * @return 	contract.ServiceOutput
 * @return 	exception.Exception
 */
// func (main *Main) CreateNew(data dataRaw.NewTransactionCreate) (contract.ServiceOutput, exception.Exception) {
// 	// create transaction
// 	transaction, exc := TransactionModel.Create(data)
// 	if exc != nil {
// 		return nil, exc
// 	}

// 	return output.NewTransactionDetail(transaction), nil
// }

// ------------------------------------------------------

/**
 * Delete Existing Transaction
 *
 * @param 	string 	code
 *
 * @return 	string
 * @return 	exception.Exception
 */
// func (main *Main) DeleteExisting(code string) (string, exception.Exception) {
// 	// delete transaction data by code
// 	transactionCode, exc := TransactionModel.Delete(code)
// 	if exc != nil {
// 		return "", exc
// 	}

// 	return transactionCode, nil
// }

// ------------------------------------------------------
