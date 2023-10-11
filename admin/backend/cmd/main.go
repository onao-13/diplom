package main

import (
	"admin/internal/app/server"
	"admin/internal/config"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	srv.Serve()
}