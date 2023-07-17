package handler

import (
	"github.com/bohoslavskyi/ip-info/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/ip-info", h.GetIPInfo)

	return router
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, model.ErrorResponse{
		Message: message,
	})
}
