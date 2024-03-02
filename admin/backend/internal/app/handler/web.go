package handler

import (
	"admin/internal/app/web"
	"fmt"
	"net/http"
)

func Page(name string, w http.ResponseWriter) {
	p, err := web.LoadPage(name)
	if err != nil {
		fmt.Fprintf(w, "Not found")
		return
	}
	fmt.Fprintf(w, string(p.Body))
}
