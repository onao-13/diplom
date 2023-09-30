package handler

import (
	"fmt"
	"net/http"

	"backend/internal/app/web"
)

func HandlePage(name string, w http.ResponseWriter) {
	p, err := web.LoadPage(name)
	if err != nil {
		fmt.Fprintf(w, "Not found")
		return
	}
	fmt.Fprintf(w, string(p.Body))
}
