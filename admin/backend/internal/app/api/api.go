package api

import (
	"admin/internal/app/controller"
	"admin/internal/app/handler"
	"admin/internal/app/middleware/security"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	AuthPanelPath = "/"
	AuthApi       = "/api/auth"
)

func Route(
	city controller.City, calls controller.ManagerCalls,
	home controller.Home, auth controller.Auth,
	authSecurity security.Auth,
	web controller.Web,
) *mux.Router {
	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return checkUserToken(next, authSecurity)
	})

	// API
	// AUTH
	r.HandleFunc("/api/auth", auth.Login).Methods(http.MethodPost)

	// CITY
	r.HandleFunc("/api/cities", city.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/cities", city.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/cities/{id:[0-9]+}", city.Preview).Methods(http.MethodGet)
	r.HandleFunc("/api/cities/{id:[0-9]+}", city.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/cities/{id:[0-9]+}", city.Delete).Methods(http.MethodDelete)
	// HOMES
	r.HandleFunc("/api/cities/{cityId:[0-9]+}/homes", home.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/cities/{cityId:[0-9]+}/homes", home.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/cities/{cityId:[0-9]+}/homes/{id:[0-9]+}", home.Preview).Methods(http.MethodGet)
	r.HandleFunc("/api/cities/{cityId:[0-9]+}/homes/{id:[0-9]+}", home.Update).Methods(http.MethodPatch)
	r.HandleFunc("/api/cities/{cityId:[0-9]+}/homes/{id:[0-9]+}", home.Delete).Methods(http.MethodDelete)
	// CALLS
	r.HandleFunc("/api/manager-calls", calls.GetAll).Methods(http.MethodGet)

	// WEB
	// file server
	fileServer := http.FileServer(http.Dir("frontend/static/"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fileServer))

	r.HandleFunc("/", web.Auth).Methods(http.MethodGet)
	r.HandleFunc("/panel", web.Panel).Methods(http.MethodGet)
	r.HandleFunc("/calls", web.Calls).Methods(http.MethodGet)
	r.HandleFunc("/cities", web.Cities).Methods(http.MethodGet)
	r.HandleFunc("/homes", web.Homes).Methods(http.MethodGet)
	r.HandleFunc("/home", web.Home).Methods(http.MethodGet)

	return r
}

func checkUserToken(next http.Handler, authSecurity security.Auth) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == AuthPanelPath ||
			r.URL.Path == AuthApi ||
			strings.HasPrefix(r.URL.Path, "/static") {
			next.ServeHTTP(w, r)
			return
		}

		token, err := r.Cookie("token")
		if err != nil {
			handler.Forbidden(w, "Токен не указан")
			return
		}

		if err := authSecurity.IsAuth(token.Value); err != nil {
			handler.Unauth(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
