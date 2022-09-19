package main

import (
	"simple-api/models"
	"simple-api/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.CreditCard{})
	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)
	r.Run()
}
