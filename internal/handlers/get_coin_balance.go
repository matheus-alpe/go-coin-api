package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/matheus-alpe/go-coin-api/api"
	"github.com/matheus-alpe/go-coin-api/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
    params := api.CoinBalanceParams{}
    decoder := schema.NewDecoder()

    err := decoder.Decode(&params, r.URL.Query()) 
    if err != nil {
	api.InternalErrorHandler(w, err)
	return
    }

    var database *tools.DatabaseInterface
    database, err = tools.NewDatabase()
    if err != nil {
	api.InternalErrorHandler(w, err)
	return
    }

    tokenDetails := (*database).GetUserCoins(params.Username)
    if tokenDetails == nil {
	api.InternalErrorHandler(w, errors.New("no coins available"))
	return
    }

    var response = api.CoinBalanceResponse{
	Balance: tokenDetails.Coins,
	Code: http.StatusOK,
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
	api.InternalErrorHandler(w, err)
	return 
    }
}
