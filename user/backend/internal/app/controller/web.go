package controller

import (
	"net/http"

	"backend/internal/app/handler"
	"backend/internal/app/web"
)

type Web struct{}

func NewWeb() Web {
	return Web{}
}

func (*Web) Index(w http.ResponseWriter, r *http.Request) {
	handler.HandlePage(web.Index, w)
}

func (*Web) Articles(w http.ResponseWriter, r *http.Request) {
	handler.HandlePage(web.Articles, w)
}

func (*Web) Locations(w http.ResponseWriter, r *http.Request) {
	handler.HandlePage(web.Locations, w)
}

func (*Web) Home(w http.ResponseWriter, r *http.Request) {
	handler.HandlePage(web.Home, w)
}
