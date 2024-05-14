package handler

import (
	"github.com/buts00/Graph/internal/database"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *database.PostgresDB
}

func NewHandler(db database.PostgresDB) Handler {
	return Handler{DB: &db}
}

func setResponseHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

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
		graphGroup.DELETE("/", h.deleteEdge)
		graphGroup.GET("/MST", h.mst)
		graphGroup.POST("/dijkstra", h.dijkstra)
	}

	return router
}
