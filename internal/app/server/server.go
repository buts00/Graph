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

//func Start(port string, db *database.PostgresDB) error {
//	router := mux.NewRouter()
//	router.HandleFunc("/", handler.HomeHandler)
//	router.HandleFunc("/graph", handler.GraphHandler(db))
//	router.HandleFunc("/graph/MST", handler.MSTHandler(db))
//	router.HandleFunc("/graph/dijkstra", handler.DijkstraHandler(db))
//
//	if err := http.ListenAndServe(port, router); err != nil {
//		return err
//	}
//	return nil
//}
