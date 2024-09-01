package cmd

import (
	"github.com/alwindoss/ab/internal/server"
)

func runServer() error {
	cfg := server.Config{
		Host:       "",
		Port:       "8080",
		DBLocation: "/home/alwin/.ab/data/ab.db",
	}
	return server.Run(&cfg)
}
