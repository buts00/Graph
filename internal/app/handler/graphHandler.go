package handlers

import (
	"github.com/buts00/Graph/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) allEdges(ctx *gin.Context) {
	curGraph, err := database.Edges(&h.DB)
	if err != nil {
		http.Error(writer, "Cannot connect to database", http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, curGraph.Edges)

}
