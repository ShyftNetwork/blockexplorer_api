package db

import (
	"github.com/ShyftNetwork/blockexplorer_api/types"
	"github.com/ShyftNetwork/blockexplorer_api/logger"
)

func returnType(str string) ([]interface{}) {
	switch {
	case str == "account":
		var account types.Account
		var accounts []types.Account
		a := []interface{}{ &account, &accounts }
		return a
	case str == "block":
		var block types.Block
		var blocks []types.Block
		b := []interface{}{ &block, &blocks }
		return b
	case str == "tx":
		var tx types.Transaction
		var txs []types.Transaction
		t := []interface{}{ &tx, &txs }
		return t
	case str == "itx":
		var itx types.InteralTransaction
		var itxs []types.InteralTransaction
		i := []interface{}{ &itx, &itxs }
		return i
	}
	return nil
}

// Query queries the postgres database and returns block data in bytes
// args[0] = query, args[1] = type, args[2] = currentPage, args[3] = pageLimit,
// args[4] = identifier
func Query(args...string) ([]byte,error) {
	d := ConnectShyftDatabase()
	result := returnType(args[1])
	switch {
	case len(args) > 4:
		limit, offset := ReturnOffset(args[2], args[3])
		if err := d.Db.Select(result[1], args[0], limit, offset, args[4]); err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil, err
		}
		payload := ReturnSerializedPayload(result[1])
		return payload, nil
	case len(args) == 4:
		limit, offset := ReturnOffset(args[2], args[3])
		if err := d.Db.Select(result[1], args[0], limit, offset); err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil, err
		}
		payload := ReturnSerializedPayload(result[1])
		return payload, nil
	case len(args) == 3:
		if err := d.Db.Get(result[0], args[0], args[2]); err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil, err
		}
		payload := ReturnSerializedPayload(result[0])
		return payload, nil
	default:
		if err := d.Db.Select(result[1], args[0]); err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil, err
		}
		payload := ReturnSerializedPayload(result[1])
		return payload, nil
	}
}

// RecordCountQuery returns count of records in specified table
func RecordCountQuery(query string) []byte {
	d := ConnectShyftDatabase()
	count := types.RecordCount{}
	if err := d.Db.Get(&count, query); err != nil {
		logger.Warn("Unable to retrieve rows: " + err.Error())
		return nil
	}
	payload := ReturnSerializedPayload(count)
	return payload
}
