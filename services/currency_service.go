package services

import (
	"context"
	"log"

	"github.com/wchopite/api-spark-v1/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CoinPrice ...
type CoinPrice struct {
	Coin           string
	Purchase_value float64
	Sale_value     float64
	Operator       string
}

//GetCurrencyBestPrice ...
func GetCurrencyBestPrice(currency string) []CoinPrice {
	var filter bson.D = bson.D{}
	var result []CoinPrice

	coll := database.GetCollection("currencyPrices")

	opts := options.Find().SetSort(bson.D{{"purchase_value", -1}, {"sale_value", -1}})

	if currency != "" {
		filter = bson.D{{"coin", currency}}
		opts.SetLimit(1)
	}

	cursor, err := coll.Find(context.Background(), filter, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var object CoinPrice
		err := cursor.Decode(&object)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, object)
	}

	return result
}

//CreateCurrencyPrice ...
func CreateCurrencyPrice(coinPrice CoinPrice) *mongo.InsertOneResult {
	coll := database.GetCollection("currencyPrices")

	result, _ := coll.InsertOne(context.Background(), coinPrice)
	return result
}
