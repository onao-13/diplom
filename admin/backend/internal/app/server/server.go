package server

import (
	"admin/internal/app/api"
	"admin/internal/app/controller"
	"admin/internal/app/middleware/database"
	"admin/internal/app/middleware/service"
	"admin/internal/config"
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg config.Config
	srv *http.Server
	ctx context.Context
}

func New(cfg config.Config) Server {
	return Server{cfg: cfg, ctx: context.Background()}
}

func (s *Server) Serve() {
	var log logrus.Logger = *logrus.New()
	
	log.Info("Сервер запускается")

	pool, err := pgxpool.New(s.ctx, s.cfg.DbUrl())
	if err != nil {
		log.Error("Ошибка подключения к базе: ", err.Error())
	}

	articleDatabase := database.NewArticle(s.ctx, pool, log)
	cityDatabase := database.NewCity(s.ctx, pool, log)

	articleService := service.NewArticle(articleDatabase)
	cityService := service.NewCity(cityDatabase)

	articleController := controller.NewArticle(articleService, log)
	cityController := controller.NewCity(cityService, log)
	locationController := controller.NewLocation()

	route := api.Route(articleController, cityController, locationController)

	s.srv = &http.Server{
		Addr: ":8085",
		Handler: route,
	}

	log.Info("Сервер запущен")

	s.srv.ListenAndServe()
}
