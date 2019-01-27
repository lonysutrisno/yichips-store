package model

import (
	// native
	"crypto/sha256"
	"encoding/base64"
	"os"
	"reflect"
	"strconv"
	"time"
)

// ------------------------------------------------------
// ------------------------------------------------------
// INTERFACE
// ------------------------------------------------------

type ApiResponseModel interface {
	AddTimestampAndSign()
}

// ------------------------------------------------------
// ------------------------------------------------------
// MAIN FUNCTIONS
// ------------------------------------------------------

/**
 * Generate Timestamp
 *
 * @return 	int
 */
func AddTimestamp() int {
	return int(time.Now().Unix())
}

// ------------------------------------------------------

/**
 * Generate Signature
 *
 * @param 	ApiResponse 	resp
 *
 * @return 	string
 */
func Sign(resp ApiResponseModel) string {
	concatedValue := ""
	v := reflect.ValueOf(resp).Elem()

	for i := 0; i < v.NumField(); i++ {
		valueType := reflect.TypeOf(v.Field(i).Interface()).Kind()
		key := v.Field(i)

		if valueType == reflect.String {

			concatedValue = concatedValue + key.String()

		} else if valueType == reflect.Int {

			keyValue := strconv.FormatInt(key.Int(), 10)
			concatedValue = concatedValue + keyValue

		}
	}

	concatedValue = concatedValue + os.Getenv("API_KEY")

	sha256Crypter := sha256.New()
	sha256Crypter.Write([]byte(concatedValue))

	return base64.URLEncoding.EncodeToString((sha256Crypter.Sum(nil)))
}
