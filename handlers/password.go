package handlers

import (
	"Password_Create_Manager/models"
	"Password_Create_Manager/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreatePassword(c *gin.Context) {
	var request struct {
		Password string `json:"password"`
		HashType string `json:"hash_type"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(request.Password, request.HashType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	password := models.Password{
		Password:  hashedPassword,
		HashType:  request.HashType,
		CreatedAt: time.Now(),
	}

	_, err = models.PasswordCollection.InsertOne(context.Background(), password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, password)
}

func GetAllPasswords(c *gin.Context) {
	var passwords []models.Password

	cursor, err := models.PasswordCollection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var password models.Password
		if err := cursor.Decode(&password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		passwords = append(passwords, password)
	}

	c.JSON(http.StatusOK, passwords)
}
