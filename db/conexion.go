package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://andersonsanta3:ybPFM91HiVIGJpBs@twittor-clouster.5h2dlo5.mongodb.net/?retryWrites=true&w=majority")

// context.TODO() -> Context default sin timeout
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con DB")
	return client
}

func haveConnection() bool {
	if err := MongoCN.Ping(context.TODO(), nil); err != nil {
		return false
	}
	return true
}
