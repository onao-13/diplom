package controller

import (
	"backend/internal/app/handler"
	"backend/internal/app/middleware/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	s service.Article
}

func NewArticle(s service.Article) Article {
	return Article{s}
}

func (a *Article) List(w http.ResponseWriter, r *http.Request) {
	res, err := a.s.List()
	if err != nil {
		handler.HandlerInternalServerError(w, "Ошибка получения статей")
		return
	}

	handler.HandleOkData(w, res)
}

func (a *Article) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID пустое")
		return
	}

	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		handler.HandlerInternalServerError(w, "Ошибка получения ID")
		return
	}

	res, err := a.s.Get(id)
	if err != nil {
		handler.HandleNotFound(w, "Статья не найдена")
		return
	}

	handler.HandleOkData(w, res)
}