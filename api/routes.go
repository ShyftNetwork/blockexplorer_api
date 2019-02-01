package api

//@NOTE Shyft setting up endpoints
import (
	"net/http"
	"github.com/ShyftNetwork/blockexplorer_api/api/blocks"
	"github.com/ShyftNetwork/blockexplorer_api/api/accounts"
	"github.com/ShyftNetwork/blockexplorer_api/api/transactions"
	"github.com/ShyftNetwork/blockexplorer_api/api/broadcast"
)

//Route stuct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes routes
type Routes []Route

// Endpoints outline all endpoints for api calls
var Endpoints = Routes{
	Route{
		"GetAllAccountLength",
		"GET",
		"/api/get_all_accounts_length",
		accounts.GetAllAccountsLength,
	},
	Route{
		"GetAccount",
		"GET",
		"/api/get_account/{address}",
		accounts.GetAccount,
	},
	Route{
		"GetAllAccounts",
		"GET",
		"/api/get_all_accounts/{currentPage}/{pageLimit}",
		accounts.GetAllAccounts,
	},
	Route{
		"GetAccountTxs",
		"GET",
		"/api/get_account_txs/{currentPage}/{pageLimit}/{address}",
		transactions.GetAccountTxs,
	},
	Route{
		"GetAllTransactionsWithoutLimit",
		"GET",
		"/api/get_all_transactions_nolimit",
		transactions.GetAllTransactionsWithoutLimit,
	},
	Route{
		"GetAllTransactionsLength",
		"GET",
		"/api/get_all_transactions_length",
		transactions.GetAllTransactionsLength,
	},
	Route{
		"GetAllTransactions",
		"GET",
		"/api/get_all_transactions/{currentPage}/{pageLimit}",
		transactions.GetAllTransactions,
	},
	Route{
		"GetTransaction",
		"GET",
		"/api/get_transaction/{txHash}",
		transactions.GetTransaction,
	},
	Route{
		Name:        "GetAllTransactionsFromBlock",
		Method:      "GET",
		Pattern:     "/api/get_all_transactions_from_block/{currentPage}/{pageLimit}/{blockNumber}",
		HandlerFunc: transactions.GetAllTransactionsFromBlock,
	},
	Route{
		"GetInternalTransactions",
		"GET",
		"/api/get_internal_transactions/{currentPage}/{pageLimit}",
		transactions.GetInternalTransactions,
	},
	Route{
		"GetInternalTransactionsByHash",
		"GET",
		"/api/get_internal_transactions/{currentPage}/{pageLimit}/{txHash}",
		transactions.GetInternalTransactionsByHash,
	},
	Route{
		"GetAllInternalTransactionsLength",
		"GET",
		"/api/get_internal_transactions_length",
		transactions.GetAllInternalTransactionsLength,
	},
	Route{
		"GetSearchQuery",
		"GET",
		"/api/search/{query}",
		transactions.GetSearchQuery,
	},
	Route{
		"GetAllBlocksWithoutLimit",
		"GET",
		"/api/get_all_blocks_nolimit",
		blocks.GetAllBlocksWithoutLimit,
	},
	Route{
		"GetAllBlocks",
		"GET",
		"/api/get_all_blocks/{currentPage}/{pageLimit}",
		blocks.GetAllBlocks,
	},
	Route{
		"GetAllBlocksLength",
		"GET",
		"/api/get_all_blocks_length",
		blocks.GetAllBlocksLength,
	},
	Route{
		"GetBlock",
		"GET",
		"/api/get_block/{blockNumber}",
		blocks.GetBlock,
	},
	Route{
		Name:        "GetRecentBlock",
		Method:      "GET",
		Pattern:     "/api/get_recent_block",
		HandlerFunc: blocks.GetRecentBlock,
	},
	Route{
		Name:        "GetAllBlocksMinedByAddress",
		Method:      "GET",
		Pattern:     "/api/get_blocks_mined/{currentPage}/{pageLimit}/{coinbase}",
		HandlerFunc: blocks.GetAllBlocksMinedByAddress,
	},
	Route{
		"Broadcast",
		"GET",
		"/api/broadcast_tx/{transaction_hash}",
		broadcast.Broadcast,
	},
}
