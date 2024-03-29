package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

const DB_NAME = "staem"
const DB_HOST = "127.0.0.1"
const DB_PORT = "5432"        //port on installation
const DB_USER = "postgres"    //default is postgres
const DB_PASS = "ShaoLiang12" //password on installation

var DB *gorm.DB

func connect() (*gorm.DB, error) {
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	return gorm.Open("postgres", "host="+DB_HOST+" port="+DB_PORT+" user="+DB_USER+" dbname="+DB_NAME+" password="+DB_PASS+" sslmode=disable")
}

func GetInstance() *gorm.DB {
	var err error

	if DB == nil {
		DB, err = connect()
	}

	if err != nil {
		panic(err)
	}

	return DB
}
