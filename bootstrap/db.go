package bootstrap

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type dbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var port, _ = strconv.Atoi(os.Getenv("DB_PORT"))

	var dbCon = dbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
	// var dbCon = dbConfig{
	// 	Host:     "127.0.0.1",
	// 	Port:     3306,
	// 	Username: "root",
	// 	Password: "",
	// 	Database: "test_golang",
	// }

	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		dbCon.Username,
		dbCon.Password,
		dbCon.Host,
		dbCon.Port,
		dbCon.Database,
	)
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	DB = db
}
