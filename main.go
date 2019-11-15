package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wchopite/api-spark-v1/config"
	"github.com/wchopite/api-spark-v1/database"
	"github.com/wchopite/api-spark-v1/handlers"
)

func main() {
	config.Init()
	c := config.GetConfig()
	database.Init()

	router := mux.NewRouter()

	router.HandleFunc("/api/currency", handlers.CreateCurrencyPrice).Methods("POST")
	router.HandleFunc("/api/currency", handlers.GetCurrencyBestPrice).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+c.Port, router))
}
