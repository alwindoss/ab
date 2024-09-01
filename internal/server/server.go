package server

import (
	"fmt"
	"net/http"

	"github.com/alwindoss/ab/pkg/addressbook"
	"go.etcd.io/bbolt"
)

type Config struct {
	Host       string
	Port       string
	DBLocation string
}

func Run(cfg *Config) error {
	db, err := bbolt.Open(cfg.DBLocation, 0600, nil)
	if err != nil {
		return err
	}
	repo := addressbook.NewBBoltRepository(db)
	svc := addressbook.NewService(repo)
	hdlr := addressbook.NewHandler(svc)

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	mux := http.NewServeMux()
	setupRoutes(mux, hdlr)

	err = http.ListenAndServe(addr, mux)
	if err != nil {
		return err
	}
	return nil
}

func setupRoutes(mux *http.ServeMux, hdlr addressbook.Handler) {
	mux.HandleFunc("POST /addressbook/v1/contacts", hdlr.CreateContacts)
}
