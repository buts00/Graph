package handler

import (
	"github.com/buts00/Graph/internal/app/graph"
	"github.com/buts00/Graph/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) allEdges(ctx *gin.Context) {
	curGraph, err := database.Edges(h.DB)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "cannot connect to db: "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, curGraph.Edges)

}

func (h *Handler) addEdge(ctx *gin.Context) {
	var myGraph graph.Graph
	edges := myGraph.Edges
	if err := ctx.BindJSON(&edges); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "failed to parse new Edge: "+err.Error())
		return
	}
	ids := make([]int, 0)
	for _, edge := range edges {
		isEdgeExist, isReversedEdgeExist, err := database.IsEdgeExist(h.DB, edge)

		if isEdgeExist || isReversedEdgeExist {
			continue
		}
		if err != nil {
			NewErrorResponse(ctx, http.StatusInternalServerError, "failed to check if edge is exists: "+err.Error())
			return
		}
		id, err := database.AddEdge(h.DB, edge)
		ids = append(ids, id)
		if err != nil {
			NewErrorResponse(ctx, http.StatusInternalServerError, "failed to add edge: "+err.Error())
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"ids": ids})

}

func (h *Handler) deleteEdge(ctx *gin.Context) {

	var myGraph graph.Graph
	edges := myGraph.Edges
	ids := make([]int, 0)
	if err := ctx.BindJSON(&edges); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "failed to parse new Edge: "+err.Error())
		return
	}

	for _, edge := range edges {
		isEdgeExist, isReversedEdgeExist, err := database.IsEdgeExist(h.DB, edge)
		if err != nil {
			NewErrorResponse(ctx, http.StatusInternalServerError, "failed to check if edge is exists: "+err.Error())
			return
		}

		if isEdgeExist {
			id, err := database.DeleteEdge(h.DB, edge)
			ids = append(ids, id)
			if err != nil {
				NewErrorResponse(ctx, http.StatusInternalServerError, "failed to delete edge: "+err.Error())
				return
			}
		} else if isReversedEdgeExist {
			id, err := database.DeleteEdge(h.DB, edge)
			ids = append(ids, id)
			if err != nil {
				NewErrorResponse(ctx, http.StatusInternalServerError, "failed to delete edge: "+err.Error())
				return
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"ids": ids})
}
