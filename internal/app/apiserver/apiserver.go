package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/buts00/Graph/internal/database"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Homepage")
	if err != nil {
		return
	}
}

func ArrayHandler(db *database.PostgresDB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		graph, err := database.Edges(db)
		if err != nil {
			http.Error(writer, " Cannot connect to database", http.StatusInternalServerError)
		}

		jsonData, err := json.Marshal(graph)
		if err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(jsonData)
		if err != nil {
			http.Error(writer, "Failed to write JSON response", http.StatusInternalServerError)
			return
		}

	}
}

func Start(port string, db *database.PostgresDB) error {
	http.HandleFunc("/", Home)
	http.HandleFunc("/array", ArrayHandler(db))
	if err := http.ListenAndServe(port, nil); err != nil {
		return err
	}
	return nil
}
