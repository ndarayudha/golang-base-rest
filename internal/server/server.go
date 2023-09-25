package server

import (
	"net/http"
	"rest_base/config"
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
