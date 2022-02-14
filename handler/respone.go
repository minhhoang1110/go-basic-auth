package handler

import (
	"encoding/json"
	"net/http"
)

type respone struct {
	success bool
	data    interface{}
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}, objectName interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	var res respone
	res.success = status == http.StatusOK
	if objectName != nil {
		res.data = map[interface{}]interface{}{objectName: object}
	} else {
		res.data = object
	}
	json.NewEncoder(writer).Encode(object)
}
