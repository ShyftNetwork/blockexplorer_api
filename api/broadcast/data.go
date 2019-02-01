package broadcast

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"bytes"
	"io/ioutil"
	"github.com/ShyftNetwork/blockexplorer_api/logger"
)

// Broadcast broadcasts tx
func Broadcast(w http.ResponseWriter, r *http.Request) {
	// Example return result (returns tx hash):
	// {"jsonrpc":"2.0","id":1,"result":"0xafa4c62f29dbf16bbfac4eea7cbd001a9aa95c59974043a17f863172f8208029"}

	// http params
	vars := mux.Vars(r)
	transactionHash := vars["transaction_hash"]

	// format the transactionHash into a proper sendRawTransaction jsonrpc request
	formattedJSON := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["%s"],"id":0}`, transactionHash))

	// send json rpc request
	resp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer(formattedJSON))
	if err != nil {
		logger.Warn("Error: " + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warn("Error: " + err.Error())
	}
	byt := []byte(string(body))

	// read json and return result as http response, be it an error or tx hash
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	txHash := dat["result"]
	if txHash == nil {
		errMap := dat["error"].(map[string]interface{})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, "ERROR:", errMap["message"]); err != nil {
			logger.Warn("Error: " + err.Error())
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, "Transaction Hash:", txHash); err != nil {
			logger.Warn("Error: " + err.Error())
		}
	}
}
