package blocks

import (
	"net/http"

	_ "github.com/lib/pq" //github.com/lib/pq needed for sqlx transactions
	"github.com/ShyftNetwork/blockexplorer_api/logger"
	"github.com/ShyftNetwork/blockexplorer_api/db"
	"github.com/gorilla/mux"
)

//GetBlock returns contextual block data
func GetBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNumber := vars["blockNumber"]
	block, err := db.Query(db.GetBlock, "block", blockNumber)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(block))
}

// GetAllBlocksWithoutLimit returns all blocks in table
func GetAllBlocksWithoutLimit(w http.ResponseWriter, r *http.Request) {

	blocks, err := db.Query(db.GetAllBlocksNoLimit, "block")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(blocks))
}

// GetAllBlocks returns all blocks
func GetAllBlocks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]

	blocks, err := db.Query(db.GetAllBlocks, "block", currentPage, pageLimit)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(blocks))
}

// GetRecentBlock returns most recent block height
func GetRecentBlock(w http.ResponseWriter, r *http.Request) {
	block, err := db.Query(db.GetRecentBlock, "block")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(block))
}

// GetAllBlocksMinedByAddress returns all blocks mined by specific address
func GetAllBlocksMinedByAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coinbase := vars["coinbase"]
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]

	blocks, err := db.Query(db.GetAllBlocksMinedByAddress, "block", currentPage, pageLimit, coinbase)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(blocks))
}

// GetAllBlocksLength Count all rows in Blocks Table
func GetAllBlocksLength(w http.ResponseWriter, r *http.Request) {
	count := db.RecordCountQuery(db.GetBlockCount)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(count))
}
