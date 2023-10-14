package main

import (
	"backend/internal/app/server"
	"backend/internal/config"
	"context"

	"github.com/sirupsen/logrus"
)

func main() {
	var log = logrus.New()
	ctx := context.Background()
	c := config.Dev()
	s := server.New(c, ctx, *log)
	s.Serve()
}
