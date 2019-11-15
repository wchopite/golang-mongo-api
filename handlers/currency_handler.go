package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wchopite/api-spark-v1/services"
)

//GetCurrencyBestPrice ...
func GetCurrencyBestPrice(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	currency := params.Get("key")

	result := services.GetCurrencyBestPrice(currency)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}

//CreateCurrencyPrice ...
func CreateCurrencyPrice(w http.ResponseWriter, r *http.Request) {
	var coinPrice services.CoinPrice
	_ = json.NewDecoder(r.Body).Decode(&coinPrice)

	w.Header().Set("content-type", "application/json")

	result := services.CreateCurrencyPrice(coinPrice)
	json.NewEncoder(w).Encode(result)
}
