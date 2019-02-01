package accounts

import (
	"net/http"

	_ "github.com/lib/pq" //github.com/lib/pq needed for sqlx transactions
	"github.com/ShyftNetwork/blockexplorer_api/logger"
	"github.com/ShyftNetwork/blockexplorer_api/db"
	"github.com/gorilla/mux"
)

// GetAllAccountsLength Count all rows in accounts Table
func GetAllAccountsLength(w http.ResponseWriter, r *http.Request) {
	count := db.RecordCountQuery(db.GetAccountCount)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(count))
}

// GetAccount returns specific account data; balance, nonce
func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	account, err := db.Query(db.GetAccount,"account", address)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(account))
}

// GetAllAccounts returns all accounts
func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]

	accounts, err := db.Query(db.GetAllAccounts,"account", currentPage, pageLimit)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(accounts))
}