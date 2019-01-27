package random

import (
	// native
	"math/rand"
)

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Random Operation Helpers
 |	----------------------
 |
 |	This helper is used to do any operation requiring randomize (e.g)
 |	make random string by length, etc.
 |	This hash helper already customed with the logic required and
 |	standarized for any services/core used here
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

/**
 * Generate Random String with Defined Length
 *
 * @params 	int 	length
 *
 * @return 	string
 */
func String(length int) string {
	// initialize set params
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	// randomize string
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}
