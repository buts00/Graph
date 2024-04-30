package handler

import (
	"github.com/buts00/Graph/internal/app/graph/algorithms"
	"github.com/buts00/Graph/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) mst(ctx *gin.Context) {
	curGraph, err := database.Edges(h.DB)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "failed to get edges")
		return
	}
	inMST := algorithms.NewMST().FindMST(curGraph)
	ctx.JSON(http.StatusOK, inMST)
}
