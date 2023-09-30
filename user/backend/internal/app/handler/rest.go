package handler

import (
	"encoding/json"
	"net/http"
)

func HandleOk(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)

	resData, err := json.Marshal(data)
	if err != nil {
	}

	w.Write(resData)
}
