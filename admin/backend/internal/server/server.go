package server

import (
	"admin/internal/app/api"
	"admin/internal/app/controller"
	"admin/internal/app/database"
	"admin/internal/app/middleware/security"
	"admin/internal/app/service"
	"admin/internal/config"
	"context"
	"fmt"
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
	var log = logrus.New()

	log.Info("Сервер запускается")

	pool, err := pgxpool.New(s.ctx, s.cfg.DbUrl())
	if err != nil {
		panic(fmt.Sprintf("Ошибка подключения к базе: %s", err))
	}

	securityAuth := security.NewAuth(s.cfg.Root, log)

	callsDatabase := database.NewManagerCalls(s.ctx, pool)
	cityDatabase := database.NewCity(s.ctx, pool)
	homeDatabase := database.NewHome(s.ctx, pool)

	callsService := service.NewManagerCalls(callsDatabase, log)
	cityService := service.NewCity(log, cityDatabase)
	homeService := service.NewHome(homeDatabase, log)

	callsController := controller.NewManagerCalls(callsService)
	cityController := controller.NewCity(cityService)
	homeController := controller.NewHome(homeService)
	authController := controller.NewAuth(securityAuth)
	webController := controller.NewWeb()

	route := api.Route(cityController, callsController, homeController, authController, securityAuth, webController)

	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.cfg.Port),
		Handler: route,
	}

	log.Infoln("Сервер запущен на порту ", s.cfg.Port)

	err = s.srv.ListenAndServe()
	if err != nil {
		log.Panicln("Ошибка запуска сервера: ", err)
		return
	}
}
