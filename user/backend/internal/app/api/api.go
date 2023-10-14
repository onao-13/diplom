package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"backend/internal/app/controller"
)

func Route(
	webController controller.Web,
	articleController controller.Article,
	homeController controller.Home,
) *mux.Router {
	r := mux.NewRouter()
	// JSON API
	// HOME
	r.HandleFunc("/api/locations/list", homeController.List).Methods(http.MethodGet)
	r.HandleFunc("/api/locations/home/{id:[0-9]+}", homeController.Get).Methods(http.MethodGet)

	// ARTICLE
	r.HandleFunc("/api/articles/list", articleController.List).Methods(http.MethodGet)
	r.HandleFunc("/api/articles/{id:[0-9]+}", articleController.Get).Methods(http.MethodGet)

	// WEB
	r.HandleFunc("/", webController.Index).Methods(http.MethodGet)
	r.HandleFunc("/articles", webController.Articles).Methods(http.MethodGet)
	r.HandleFunc("/home", webController.Home).Methods(http.MethodGet)
	r.HandleFunc("/locations", webController.Locations).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("./../frontend/static/"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fileServer))
	return r
}
