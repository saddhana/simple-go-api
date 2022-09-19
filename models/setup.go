package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "test"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	DB = db
	return db
}
