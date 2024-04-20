package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type err struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, err{message})
}
