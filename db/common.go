package db

import (
	"encoding/json"
	"github.com/ShyftNetwork/blockexplorer_api/logger"
	"strconv"
)

// ReturnSerializedPayload returns serliazed payload to client
func ReturnSerializedPayload(b interface{}) ([]byte) {
	serializedPayload, err := json.Marshal(b)
	if err != nil {
		logger.Warn("Unable to serialize row: " + err.Error())
		return nil
	}
	return serializedPayload
}

// ReturnOffset returns pageLimit and offset both int64
func ReturnOffset(page, limit string) (int64, int64) {
	currentPage := StringToInteger(page)
	pageLimit := StringToInteger(limit)
	var offset = (currentPage - 1) * pageLimit
	return pageLimit, offset
}

// StringToInteger returns int64 from string params
func StringToInteger(str string) int64 {
	response, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logger.Warn("Error converting params: " + err.Error())
	}
	return response
}

