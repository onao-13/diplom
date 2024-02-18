package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"backend/internal/app/controller"
)

func Route(
	web controller.Web,
	city controller.City,
	article controller.Article,
	home controller.Home,
	managerCall controller.ManagerCall,
) *mux.Router {
	r := mux.NewRouter()
	// REST API
	// MANAGER CALL
	r.HandleFunc("/api/send-manager-call", managerCall.Send).Methods(http.MethodPost)

	// CITY
	r.HandleFunc("/api/cities", city.GetAll).Methods(http.MethodGet)

	// HOME
	r.HandleFunc("/api/homes/{id:[0-9]+}", home.Get).Methods(http.MethodGet)

	// ARTICLE
	r.HandleFunc("/api/articles", article.List).Methods(http.MethodGet)
	r.HandleFunc("/api/articles/{id:[0-9]+}", article.Get).Methods(http.MethodGet)

	// WEB
	r.HandleFunc("/", web.Index).Methods(http.MethodGet)
	r.HandleFunc("/articles", web.Articles).Methods(http.MethodGet)
	r.HandleFunc("/home", web.Home).Methods(http.MethodGet)
	r.HandleFunc("/locations", web.Locations).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("frontend/static/"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fileServer))
	return r
}
