package utils

import (
	"admin/internal/app/handler"
	"fmt"
	"net/http"
)

func GetHeader(w http.ResponseWriter, r *http.Request, header string) string {
	ids := r.Header.Get(header)
	if len(ids) == 0 {
		handler.HandleBadRequest(w, fmt.Sprintf("%s value is empty", header))
		return ""
	}

	return ids
}