package main

import (
	"backend/internal/app/server"
	"backend/internal/config"
)

func main() {
	c := config.Dev()
	s := server.New(c)
	s.Serve()
}
