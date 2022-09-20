package models

import (
	"github.com/jinzhu/gorm"
)

type CreditCard struct {
	gorm.Model
	Number string
	// This is for mandatory
	// UserID uint
	// This is for allow nil
	UserID *uint
	User   *User
}
