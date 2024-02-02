package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/buts00/Graph/internal/app/graph"
	"github.com/buts00/Graph/internal/app/graph/Algorithms"
	"github.com/buts00/Graph/internal/database"
	"github.com/gorilla/mux"
	"io"
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

func handlePostRequest(writer http.ResponseWriter, request *http.Request, db *database.PostgresDB) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var newEdge graph.Edge
	if err = json.Unmarshal(body, &newEdge); err != nil {
		http.Error(writer, "Failed to unmarshal JSON data", http.StatusBadRequest)
		return
	}

	if err = database.AddEdge(db, newEdge.Source, newEdge.Destination, newEdge.Weight); err != nil {
		http.Error(writer, "Failed to add edge", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func handleGetRequest(writer http.ResponseWriter, db *database.PostgresDB) {
	curGraph, err := database.Edges(db)
	if err != nil {
		http.Error(writer, "Cannot connect to database", http.StatusInternalServerError)
		return
	}

	idInMst := Algorithms.NewMST().FindMST(curGraph)
	data := DataToLoad{curGraph.Edges, idInMst}
	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(writer, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}

	_, err = writer.Write(jsonData)

	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GraphHandler(db *database.PostgresDB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Content-Type", "application/json")

		if request.Method == http.MethodPost {
			handlePostRequest(writer, request, db)
		}
		handleGetRequest(writer, db)

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
