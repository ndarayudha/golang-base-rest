package server

import (
	"golang_restfull_api/config"
	"net/http"
	"time"
)

func NewServer(config *config.Config, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         config.Server.Port,
		ReadTimeout:  time.Second * config.Server.ReadTimeout,
		WriteTimeout: time.Second * config.Server.WriteTimeout,
		Handler:      router,
	}
}
