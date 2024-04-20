package apiserver

import (
	"github.com/buts00/Graph/internal/app/handlers"
	"github.com/buts00/Graph/internal/database"
	"github.com/gorilla/mux"
	"net/http"
)

func Start(port string, db *database.PostgresDB) error {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/graph", handlers.GraphHandler(db))
	router.HandleFunc("/graph/MST", handlers.MSTHandler(db))
	router.HandleFunc("/graph/dijkstra", handlers.DijkstraHandler(db))

	if err := http.ListenAndServe(port, router); err != nil {
		return err
	}
	return nil
}
