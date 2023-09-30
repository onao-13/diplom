package controller

import (
	"net/http"
	"os"
)

type Article struct {
	
}

func NewArticle() Article {
	return Article{}
}

func (*Article) Get(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./json/articles.json")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}