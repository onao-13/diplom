package controller

import (
	"admin/internal/app/handler"
	"admin/internal/app/web"
	"net/http"
)

type Web struct {
}

func NewWeb() Web {
	return Web{}
}

func (Web) Panel(w http.ResponseWriter, r *http.Request) {
	handler.Page(web.Panel, w)
}

func (Web) Auth(w http.ResponseWriter, r *http.Request) {
	handler.Page(web.Auth, w)
}

func (Web) Calls(w http.ResponseWriter, r *http.Request) {
	handler.Page(web.Calls, w)
}

func (Web) Cities(w http.ResponseWriter, r *http.Request) {
	handler.Page(web.Cities, w)
}

func (Web) Homes(w http.ResponseWriter, r *http.Request) {
	handler.Page(web.Homes, w)
}

func (Web) Home(w http.ResponseWriter, r *http.Request) {
	handler.Page(web.Home, w)
}
