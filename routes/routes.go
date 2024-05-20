package routes

import (
	"Password_Create_Manager/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/create-password", handlers.CreatePassword)
	r.GET("/passwords", handlers.GetAllPasswords)

	return r
}
