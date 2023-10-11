package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"backend/internal/app/controller"
)

func Route(
	locationController controller.Location, 
	webController controller.Web,
	articleController controller.Article,
	homeController controller.Home,
) *mux.Router {
	r := mux.NewRouter()
	// JSON API
	r.HandleFunc("/api/locations/new", locationController.Get).Methods(http.MethodGet)
	r.HandleFunc("/api/articles/new", articleController.Get).Methods(http.MethodGet)
	r.HandleFunc("/api/locations/home/{id: [0-9]+}", homeController.Data).Methods(http.MethodGet)

	// WEB
	r.HandleFunc("/", webController.Index).Methods(http.MethodGet)
	r.HandleFunc("/articles", webController.Articles).Methods(http.MethodGet)
	r.HandleFunc("/home", webController.Home).Methods(http.MethodGet)
	r.HandleFunc("/locations", webController.Locations).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("./../frontend/static/"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fileServer))
	return r
}
