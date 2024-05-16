package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func Run(port string, handler http.Handler) error {

	server := Server{
		httpServer: &http.Server{
			Addr:           port,
			MaxHeaderBytes: 1 << 20,
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}

	return server.httpServer.ListenAndServe()
}
