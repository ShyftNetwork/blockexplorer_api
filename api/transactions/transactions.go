package transactions

import (
	"net/http"
	_ "github.com/lib/pq" //github.com/lib/pq needed for sqlx transactions
	"github.com/ShyftNetwork/blockexplorer_api/logger"
	"github.com/ShyftNetwork/blockexplorer_api/db"
	"github.com/gorilla/mux"
)

// GetAllTransactionsLength Count all rows in Blocks Table
func GetAllTransactionsLength(w http.ResponseWriter, r *http.Request) {
	count := db.RecordCountQuery(db.GetTransactionCount)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(count))
}

// GetAllTransactionsWithoutLimit returns all rows in Blocks Table
func GetAllTransactionsWithoutLimit(w http.ResponseWriter, r *http.Request) {
	txs, err := db.Query(db.GetAllTransactionsNoLimit, "tx")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(txs))
}

// GetTransaction gets txs
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txHash := vars["txHash"]
	transaction, err := db.Query(db.GetTransaction, "tx", txHash)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(transaction))
}

// GetAllTransactions gets txs
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]

	txs, err := db.Query(db.GetAllTransactions, "tx", currentPage, pageLimit)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(txs))
}

// GetAllTransactionsFromBlock returns all txs from specified block
func GetAllTransactionsFromBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]
	blockNumber := vars["blockNumber"]

	transactions, err := db.Query(db.GetAllTransactionsFromBlock, "tx", currentPage, pageLimit, blockNumber)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(transactions))
}

// GetAccountTxs returns account txs
func GetAccountTxs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]

	transactions, err := db.Query(db.GetAccountTransactions, "tx", currentPage, pageLimit, address)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(transactions))
}

// GetAllInternalTransactionsLength Count all rows in Blocks Table
func GetAllInternalTransactionsLength(w http.ResponseWriter, r *http.Request) {
	count := db.RecordCountQuery(db.GetInternalTransactionLength)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(count))
}

//GetInternalTransactionsByHash gets internal txs specified by hash
func GetInternalTransactionsByHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txHash := vars["txHash"]
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]

	transactions, err := db.Query(db.GetInternalTransaction, "itx", currentPage, pageLimit, txHash)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(transactions))
}

//GetInternalTransactions gets internal txs
func GetInternalTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]

	transactions, err := db.Query(db.GetAllInternalTransactions, "itx",currentPage, pageLimit)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(transactions))
}


// GetSearchQuery returns search query from tx table
func GetSearchQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]

	response, err := db.Query(db.SearchQuery, query, "tx")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(response))
}

