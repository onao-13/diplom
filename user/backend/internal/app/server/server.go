package server

import (
	"context"
	"fmt"
	"net/http"

	"backend/internal/app/api"
	"backend/internal/app/controller"
	"backend/internal/app/middleware/database"
	"backend/internal/app/middleware/service"
	"backend/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg config.Config
	srv http.Server
	ctx context.Context
	log logrus.Logger
}

func New(cfg config.Config, ctx context.Context, log logrus.Logger) Server {
	return Server{cfg: cfg, ctx: ctx, log: log}
}

func (s *Server) Serve() {
	pool, err := pgxpool.New(s.ctx, s.cfg.DbUrl())
	if err != nil {
		panic(fmt.Sprintf("Ошибка подключения к базе данных", err.Error()))
	}

	s.log.Infoln("Сервер запускается")

	articleDatabase := database.NewArticle(pool, s.ctx)
	homeDatabase := database.NewHome(s.ctx, pool)

	articleService := service.NewArticle(articleDatabase, s.log)
	homeService := service.NewHome(homeDatabase, s.log)

	webController := controller.NewWeb()
	articleController := controller.NewArticle(articleService)
	homeController := controller.NewHome(homeService, s.log)

	route := api.Route(webController, articleController, homeController)

	s.srv = http.Server{
		Addr:    fmt.Sprintf(":%s", s.cfg.Port),
		Handler: route,
	}

	s.log.Infoln("Сервер запущен")

	s.srv.ListenAndServe()
}
