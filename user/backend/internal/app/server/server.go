package server

import (
	"backend/internal/app/database"
	"backend/internal/app/service"
	"context"
	"fmt"
	"github.com/rs/cors"
	"net/http"

	"backend/internal/app/api"
	"backend/internal/app/controller"
	"backend/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg config.Config
	srv http.Server
	ctx context.Context
	log *logrus.Logger
}

func New(cfg config.Config, ctx context.Context, log *logrus.Logger) Server {
	return Server{cfg: cfg, ctx: ctx, log: log}
}

func (s *Server) Serve() {
	pool, err := pgxpool.New(s.ctx, s.cfg.DbUrl())
	if err != nil {
		panic(fmt.Sprintf("Ошибка подключения к базе данных: %s", err.Error()))
	}

	s.log.Infoln("Сервер запускается")

	articleDatabase := database.NewArticle(pool, s.ctx)
	homeDatabase := database.NewHome(s.ctx, pool)
	cityDatabase := database.NewCity(s.ctx, pool)
	managerCallDatabase := database.NewManagerCall(s.ctx, pool)

	articleService := service.NewArticle(articleDatabase, s.log)
	homeService := service.NewHome(homeDatabase, s.log)
	cityService := service.NewCity(cityDatabase, s.log)
	managerCallService := service.NewManagerCall(s.log, managerCallDatabase)

	webController := controller.NewWeb()
	articleController := controller.NewArticle(articleService)
	cityController := controller.NewCity(cityService)
	homeController := controller.NewHome(homeService)
	managerCallController := controller.NewManagerCall(managerCallService)

	route := api.Route(
		webController,
		cityController,
		articleController,
		homeController,
		managerCallController,
	)

	corsOpt := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	s.log.Infoln("Сервер запущен на порту: ", s.cfg.Port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", s.cfg.Port), corsOpt.Handler(route))
	if err != nil {
		s.log.Panicf("Ошибка запуска сервера: %s", err)
	}
}
