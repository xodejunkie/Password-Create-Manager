package main

import (
	"Password_Create_Manager/models"
	"Password_Create_Manager/routes"
)

func main() {
	models.ConnectDB()

	r := routes.InitRoutes()
	r.Run(":8080")
}
