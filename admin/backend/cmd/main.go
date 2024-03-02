package main

import (
	"admin/internal/config"
	"admin/internal/server"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	srv.Serve()
}
