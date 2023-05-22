package db

import (
	"context"
	"time"

	"github.com/John-Santa/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ExistUser(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoCN.Database("twittor")
	colection := database.Collection("usuarios")

	condicion := bson.M{"email": email}

	var result models.Usuario
	err := colection.FindOne(ctx, condicion).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
