package api

import (
	"admin/internal/app/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Route(article controller.Article, city controller.City, location controller.Location) *mux.Router {
	r := mux.NewRouter()
	// API
	// ARTICLES
	r.HandleFunc("/api/article/create", article.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/article/{id: [0-9]+}/preview", article.Preview).Methods(http.MethodGet)
	r.HandleFunc("/api/article/{id: [0-9]+}/update", article.Update).Methods(http.MethodPatch)
	r.HandleFunc("/api/article/{id: [0-9]+}/delete", article.Delete).Methods(http.MethodDelete)
	// CITY
	r.HandleFunc("/api/city/create", city.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/city/{id: [0-9]+}/preview", city.Preview).Methods(http.MethodGet)
	r.HandleFunc("/api/city/{id: [0-9]+}/update", city.Update).Methods(http.MethodPatch)
	r.HandleFunc("/api/city/{id: [0-9]+}/delete", city.Delete).Methods(http.MethodDelete)
	// HOMES
	r.HandleFunc("/api/location/create", location.Create).Methods(http.MethodGet)
	r.HandleFunc("/api/location/{id: [0-9]+}/preview", location.Preview).Methods(http.MethodGet)
	r.HandleFunc("/api/location/{id: [0-9]+}/update", location.Update).Methods(http.MethodPatch)
	r.HandleFunc("/api/location/{id: [0-9]+}/delete", location.Delete).Methods(http.MethodDelete)

	// WEB
	return r
}