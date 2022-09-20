package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string
	FirstName   string
	LastName    string
	CreditCards []CreditCard
}

func (user User) FullName() string {
	return user.FirstName + " " + user.LastName
}
