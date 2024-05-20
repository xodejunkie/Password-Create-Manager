package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var PasswordCollection *mongo.Collection

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	PasswordCollection = client.Database("passwordmanager").Collection("passwords")
}

type Password struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Password  string    `json:"password" bson:"password"`
	HashType  string    `json:"hash_type" bson:"hash_type"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
