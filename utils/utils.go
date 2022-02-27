package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(rw http.ResponseWriter, data interface{}) {
	json.NewEncoder(rw).Encode(data)
}

func ErrorResponse(data interface{}) map[string]interface{} {
	resp := map[string]interface{}{"status": false}
	if data != nil {
		resp["result"] = data
	}
	return resp
}

func SuccessResponse(data interface{}) map[string]interface{} {
	resp := map[string]interface{}{"status": true}
	if data != nil {
		resp["result"] = data
	}
	return resp
}
