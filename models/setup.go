package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(Task{})

	db.AutoMigrate(User{})

	db.AutoMigrate(CreditCard{})
	db.Model(CreditCard{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	db.AutoMigrate(Figure{})
	db.AutoMigrate(Square{})
	db.AutoMigrate(Triangle{})
}
