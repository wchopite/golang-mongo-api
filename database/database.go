package database

import (
	"context"
	"log"

	"github.com/wchopite/api-spark-v1/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

//Init ...
func Init() {
	c := config.GetConfig()

	clientOptions := options.Client().ApplyURI(c.DbStringConn)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	log.Println("Connected to database")
	mongoClient = client
}

//GetDB ...
func GetDB() *mongo.Client {
	return mongoClient
}

//GetCollection ...
func GetCollection(coll string) *mongo.Collection {
	c := config.GetConfig()
	collection := mongoClient.Database(c.DbName).Collection(coll)

	return collection
}
