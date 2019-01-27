package model

import (
	"yichips/bootstrap"
)

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
 * Init
 */
func init() {
	// migrateDown()
	// migrateUp()

}

// ------------------------------------------------------

/**
 * Migrate Up DB
 */
func migrateUp() {
	DB.CreateTable(&Transaction{}, &TransactionDetail{})
}

// ------------------------------------------------------

/**
 * Migrate Down DB
 */
func migrateDown() {
	DB.DropTable(&User{}, &Product{}, &Cart{})
}
