package server

import (
	"fmt"
	"net/http"

	"backend/internal/app/api"
	"backend/internal/app/controller"
	"backend/internal/config"
)

type Server struct {
	cfg config.Config
	srv http.Server
}

func New(cfg config.Config) Server {
	return Server{cfg: cfg}
}

func (s *Server) Serve() {
	locationController := controller.NewLocation()
	webController := controller.NewWeb()
	articleController := controller.NewArticle()
	homeController := controller.NewHome()

	route := api.Route(locationController, webController, articleController, homeController)

	s.srv = http.Server{
		Addr:    fmt.Sprintf(":%s", s.cfg.Port),
		Handler: route,
	}

	fmt.Println("Server is starting")

	s.srv.ListenAndServe()
}
