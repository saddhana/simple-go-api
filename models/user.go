package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	FirstName   string
	LastName    string
	CreditCards []CreditCard
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.Name == "" {
		tx.Statement.SetColumn("Name", user.FullName())
		fmt.Println(">>>>  it's work and firing....")
	}
	return
}

func (user User) FullName() string {
	return user.FirstName + " " + user.LastName
}
