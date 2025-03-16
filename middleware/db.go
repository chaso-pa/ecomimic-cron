package middleware

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type DataBaseHandler struct {
	Db  *gorm.DB
	Err error
}

type DBConfig struct {
	Name     string
	Pass     string
	Addr     string
	Database string
}

var dbHandler *DataBaseHandler
var cryptonDbHandler *DataBaseHandler

func ConDb() {
	db, err := gorm.Open(mysql.Open(os.Getenv(`DATABASE_URL`)), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting database: %v", err)
	}
	dbHandler = &DataBaseHandler{db, err}
}

func ConCryptonDb() {
	db, err := gorm.Open(mysql.Open(os.Getenv(`CRYPTON_DATABASE_URL`)), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting crypton database: %v", err)
	}
	cryptonDbHandler = &DataBaseHandler{db, err}
}

func GetDb() *gorm.DB {
	return dbHandler.Db
}

func GetCryptonDb() *gorm.DB {
	return cryptonDbHandler.Db
}
