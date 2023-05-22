package db

import (
	"context"
	"time"

	"github.com/John-Santa/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
