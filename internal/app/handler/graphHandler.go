package handler

import (
	"github.com/buts00/Graph/internal/app/graph"
	"github.com/buts00/Graph/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) allEdges(ctx *gin.Context) {
	curGraph, err := database.Edges(&h.DB)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "cannot connect to db: "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, curGraph.Edges)

}

type Test struct {
	s string
}

func (h *Handler) addEdge(ctx *gin.Context) {

	var newEdge graph.Edge
	if err := ctx.BindJSON(&newEdge); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	logrus.Info(newEdge, "1221")
	ctx.JSON(http.StatusOK, Test{s: "1"})
	ctx.Status(http.StatusOK)

}
