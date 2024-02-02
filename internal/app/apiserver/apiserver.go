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
func ReverseEdge(edge graph.Edge) graph.Edge {
	temp := edge.Source
	edge.Source = edge.Destination
	edge.Destination = temp
	return edge
}

func processEdge(db *database.PostgresDB, newEdge graph.Edge) error {
	isEdgeExist, err := database.IsEdgeExist(db, newEdge)
	if err != nil {
		return fmt.Errorf("failed to check edge existence: %v", err)
	}

	reversedNewEdge := ReverseEdge(newEdge)
	isReversedEdge, err := database.IsEdgeExist(db, reversedNewEdge)
	if err != nil {
		return fmt.Errorf("failed to check reversed edge existence: %v", err)
	}

	if isEdgeExist {
		if err := database.DeleteEdge(db, newEdge); err != nil {
			return fmt.Errorf("failed to remove edge: %v", err)
		}
	} else if isReversedEdge {
		if err := database.DeleteEdge(db, reversedNewEdge); err != nil {
			return fmt.Errorf("failed to remove reversed edge: %v", err)
		}
	} else {
		if err := database.AddEdge(db, newEdge); err != nil {
			return fmt.Errorf("failed to add edge: %v", err)
		}
	}

	return nil
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

	if err := processEdge(db, newEdge); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
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
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
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
