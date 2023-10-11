package controller

import (
	"admin/internal/app/handler"
	"admin/internal/app/middleware/service"
	"admin/internal/payload"
	"admin/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Article struct {
	s   service.Article
	log logrus.Logger
}

func NewArticle(s service.Article, log logrus.Logger) Article {
	return Article{s: s, log: log}
}

func (a *Article) Create(w http.ResponseWriter, r *http.Request) {
	var article payload.Article
	var err error

	if err = json.NewDecoder(r.Body).Decode(&article); err != nil {
		handler.HandleDecodeJsonError(w, err)
		return
	}

	if err = a.s.Create(article); err != nil {
		handler.HandlerInternalServerError(w, "Error create article. Try again")
		return
	}

	handler.HandleOkMsg(w, "")
}

func (a *Article) Preview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ids, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID is empty")
		return
	}

	id := utils.ParseInt(w, ids)

	art, err := a.s.Preview(id)
	if err != nil {
		handler.HandleNotFound(w, "")
		return
	}

	data := map[string]interface{}{
		"article": art,
	}

	handler.HandleOkData(w, data)
}

func (a *Article) Update(w http.ResponseWriter, r *http.Request) {
	var art payload.Article

	vars := mux.Vars(r)
	ids, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID is empty")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		handler.HandleDecodeJsonError(w, err)
		return
	}
	
	id := utils.ParseInt(w, ids)
	
	if err := a.s.Update(id, art); err != nil {
		handler.HandlerInternalServerError(w, "Error update article. Try again")
		return
	}

	handler.HandleOkMsg(w, "")
}

func (a *Article) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID is empty")
		return
	}

	id := utils.ParseInt(w, ids)

	if err := a.s.Delete(id); err != nil {
		handler.HandlerInternalServerError(w, "Error delete article. Try again")
		return
	}

	handler.HandleOkMsg(w, "")
}
