package hash

import (
	// native
	"crypto/sha256"
	"encoding/base64"
	"time"

	// yichips
	"yichips/utils/helper/random"
)

/*
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
 |
 |	Hash Operation Helpers
 |	----------------------
 |
 |	This helper is used to do any operation requiring hashing (e.g)
 |	password hash make, etc.
 |	This hash helper already customed with the logic required and
 |	standarized for any services/core used here
 |
 |--------------------------------------------------------------------
 |--------------------------------------------------------------------
*/

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Make hash
 *
 * @param 	feed 	string
 *
 * @return 	string
 */
func Make(feed string) string {
	// hash feed
	sha256Crypter := sha256.New()
	sha256Crypter.Write([]byte(feed))

	return base64.URLEncoding.EncodeToString((sha256Crypter.Sum(nil)))
}

// ------------------------------------------------------

/**
 * Check/compare feed with existing hashed challenge
 *
 * @param 	feed 		string
 * @param 	challenge  	string
 *
 * @return 	bool
 */
func Check(feed string, challenge string) bool {
	// hash feed
	hashedFeed := Make(feed)

	// compare hased feed with the challenge
	if hashedFeed != challenge {
		return false
	}

	return true
}

// ------------------------------------------------------

/**
 * Generate Randomized Hashed String
 *
 * @return string
 */
func Randomize() string {
	// generate random string
	word := random.String(5)

	return Make(word + time.Now().Format("20060102150405"))
}
