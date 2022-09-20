package main

import (
	"simple-api/models"
	"simple-api/routes"
)

func main() {

	db := models.SetupDB()
	models.MigrateDB(db)

	r := routes.SetupRoutes(db)
	r.Run()
}
