package handler

import (
	"github.com/buts00/Graph/internal/app/graph/algorithms"
	"github.com/buts00/Graph/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) dijkstra(ctx *gin.Context) {

	startPoint, err := strconv.Atoi(ctx.Query("s"))
	endPoint, err1 := strconv.Atoi(ctx.Query("d"))
	if err != nil || err1 != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Parameter 'source' and 'destination' must be an integer")
		return
	}

	curGraph, err := database.Edges(h.DB)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "cannot connect to db: "+err.Error())
		return
	}

	isSourceExists, isDestinationExists := false, false
	for _, edge := range curGraph.Edges {
		if *edge.Source == startPoint || *edge.Destination == startPoint {
			isSourceExists = true
		}
		if *edge.Source == endPoint || *edge.Destination == endPoint {
			isDestinationExists = true
		}
		if isDestinationExists && isSourceExists {
			break
		}
	}

	if !isSourceExists || !isDestinationExists {
		NewErrorResponse(ctx, http.StatusNotFound, "node not found")
		return
	}

	path, distance := algorithms.NewDijkstra().FindDijkstra(startPoint, endPoint, curGraph)

	ctx.JSON(http.StatusOK, gin.H{
		"distance": distance,
		"path":     path,
	})
}
