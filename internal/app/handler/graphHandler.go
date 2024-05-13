package handler

import (
	"fmt"
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

	for _, edge := range edges {

		if isEdgeExist, err := database.IsEdgeExist(h.DB, edge); err != nil || isEdgeExist {
			if isEdgeExist {
				NewErrorResponse(ctx, http.StatusBadRequest, "the edge is already exists")
			}

			if err != nil {
				NewErrorResponse(ctx, http.StatusInternalServerError, "failed to check if edge is exists: "+err.Error())
			}

			return

		}
	}

	ids := make([]int, 0)
	for _, edge := range edges {
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
		if isEdgeExist, err := database.IsEdgeExist(h.DB, edge); err != nil || !isEdgeExist {
			if isEdgeExist {
				NewErrorResponse(ctx, http.StatusBadRequest, "the edge isn't exists")
			}
			if err != nil {
				NewErrorResponse(ctx, http.StatusInternalServerError, "failed to check if edge is exists: "+err.Error())
			}

			return

		}
	}

	for _, edge := range edges {
		id, err := database.DeleteEdge(h.DB, edge)
		if err != nil {
			edge.Source, edge.Destination = edge.Destination, edge.Source
			id, err = database.DeleteEdge(h.DB, edge)
			if err != nil {
				NewErrorResponse(ctx, http.StatusInternalServerError, "failed to delete edge: "+err.Error())
				return
			}
		}
		ids = append(ids, id)
	}
	fmt.Println(ids)
	ctx.JSON(http.StatusOK, gin.H{"ids": ids})
}
