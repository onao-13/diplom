package utils

import (
	"admin/internal/app/handler"
	"net/http"
	"strconv"
)

func ParseInt(w http.ResponseWriter, parseP string) int64 {
	res, err := strconv.ParseInt(parseP, 10, 64)
	if err != nil {
		handler.HandlerInternalServerError(w, "Error get ID")
		return 0
	}

	return res
}
