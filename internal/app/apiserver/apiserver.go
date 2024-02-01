package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/buts00/Graph/internal/app/graph"
	"github.com/buts00/Graph/internal/app/graph/Algorithms"
	"github.com/buts00/Graph/internal/database"
	"github.com/gorilla/mux"
	"net/http"
)

type DataToLoad struct {
	Edges []graph.Edge
	InMst []int
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Homepage")
	if err != nil {
		return
	}
}

func GraphHandler(db *database.PostgresDB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Content-Type", "application/json")

		curGraph, err := database.Edges(db)

		if err != nil {
			http.Error(writer, "Cannot connect to database", http.StatusInternalServerError)
		}

		idInMst := Algorithms.NewMST().FindMST(curGraph)
		data := DataToLoad{curGraph.Edges, idInMst}
		jsonData, err := json.Marshal(data)
		_, err = writer.Write(jsonData)

		if err != nil {
			http.Error(writer, "Failed to write JSON response", http.StatusInternalServerError)
			return
		}

		if err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	}
}

func Start(port string, db *database.PostgresDB) error {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/graph", GraphHandler(db))

	if err := http.ListenAndServe(port, router); err != nil {
		return err
	}

	return nil
}
