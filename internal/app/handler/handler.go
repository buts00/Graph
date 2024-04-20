package handler

import (
	"github.com/buts00/Graph/internal/database"
	"github.com/gin-gonic/gin"
)

//func HomeHandler(w http.ResponseWriter, r *http.Request) {
//	_, err := fmt.Fprintf(w, "Homepage")
//	if err != nil {
//		return
//	}
//}
//func reverseEdge(edge graph.Edge) graph.Edge {
//	temp := edge.Source
//	edge.Source = edge.Destination
//	edge.Destination = temp
//	return edge
//}
//
//func processEdge(db *database.PostgresDB, newEdge graph.Edge) error {
//	isEdgeExist, err := database.IsEdgeExist(db, newEdge)
//	if err != nil {
//		return fmt.Errorf("failed to check edge existence: %v", err)
//	}
//
//	reversedNewEdge := reverseEdge(newEdge)
//	isReversedEdge, err := database.IsEdgeExist(db, reversedNewEdge)
//	if err != nil {
//		return fmt.Errorf("failed to check reversed edge existence: %v", err)
//	}
//
//	if isEdgeExist {
//		if err := database.DeleteEdge(db, newEdge); err != nil {
//			return fmt.Errorf("failed to remove edge: %v", err)
//		}
//	} else if isReversedEdge {
//		if err := database.DeleteEdge(db, reversedNewEdge); err != nil {
//			return fmt.Errorf("failed to remove reversed edge: %v", err)
//		}
//	} else {
//		if err := database.AddEdge(db, newEdge); err != nil {
//			return fmt.Errorf("failed to add edge: %v", err)
//		}
//	}
//
//	return nil
//}
//
//func writeJSONResponse(writer http.ResponseWriter, data interface{}) {
//	jsonData, err := json.Marshal(data)
//	if err != nil {
//		http.Error(writer, "Failed to write JSON response", http.StatusInternalServerError)
//		return
//	}
//
//	_, err = writer.Write(jsonData)
//	if err != nil {
//		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//}
//
//func handleGraphPostRequest(writer http.ResponseWriter, request *http.Request, db *database.PostgresDB) {
//	body, err := io.ReadAll(request.Body)
//	if err != nil {
//		http.Error(writer, "Failed to read request body", http.StatusBadRequest)
//		return
//	}
//
//	var newEdge graph.Edge
//	if err = json.Unmarshal(body, &newEdge); err != nil {
//		http.Error(writer, "Failed to unmarshal JSON data", http.StatusBadRequest)
//		return
//	}
//
//	if err := processEdge(db, newEdge); err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	writer.WriteHeader(http.StatusOK)
//}
//
//func handleGraphGetRequest(writer http.ResponseWriter, db *database.PostgresDB) {
//	curGraph, err := database.Edges(db)
//	if err != nil {
//		http.Error(writer, "Cannot connect to database", http.StatusInternalServerError)
//		return
//	}
//	writeJSONResponse(writer, curGraph.Edges)
//}

//func setResponseHeaders(writer http.ResponseWriter) {
//	writer.Header().Set("Access-Control-Allow-Origin", "*")
//	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
//	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//	writer.Header().Set("Content-Type", "application/json")
//}

//func GraphHandler(db *database.PostgresDB) http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		setResponseHeaders(writer)
//		if request.Method == http.MethodPost {
//			handleGraphPostRequest(writer, request, db)
//		}
//
//		handleGraphGetRequest(writer, db)
//
//	}
//}
//
//func MSTHandler(db *database.PostgresDB) http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		setResponseHeaders(writer)
//
//		curGraph, err := database.Edges(db)
//		if err != nil {
//			http.Error(writer, "Cannot connect to database", http.StatusInternalServerError)
//			return
//		}
//
//		inMST := Algorithms.NewMST().FindMST(curGraph)
//		writeJSONResponse(writer, inMST)
//
//	}
//}
//
//func handleDijkstraGetRequest(writer http.ResponseWriter, db *database.PostgresDB, startPoint int) {
//	curGraph, err := database.Edges(db)
//	if err != nil {
//		http.Error(writer, "Cannot connect to database", http.StatusInternalServerError)
//		return
//	}
//	distance := dijkstra.NewDijkstra().FindDijkstra(startPoint, curGraph)
//	writeJSONResponse(writer, distance)
//}
//
//func handleDijkstraPostRequest(writer http.ResponseWriter, request *http.Request) int {
//	body, err := io.ReadAll(request.Body)
//	if err != nil {
//		http.Error(writer, "Failed to unmarshal JSON data", http.StatusBadRequest)
//		return 0
//	}
//
//	var startPoint int
//	if err = json.Unmarshal(body, &startPoint); err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return 0
//	}
//	writer.WriteHeader(http.StatusOK)
//	return startPoint
//}
//
//func DijkstraHandler(db *database.PostgresDB) http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		setResponseHeaders(writer)
//		if request.Method == http.MethodPost {
//			startPoint := handleDijkstraPostRequest(writer, request)
//			handleDijkstraGetRequest(writer, db, startPoint)
//		}
//
//	}
//}

type Handler struct {
	DB database.PostgresDB
}

func NewHandler(db database.PostgresDB) Handler {
	return Handler{DB: db}
}

func setResponseHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(setResponseHeadersMiddleware())
	graphGroup := router.Group("/graph")
	{
		graphGroup.GET("/", h.allEdges)
		graphGroup.POST("/", h.addEdge)
	}

	return router
}
