package handler

import (
	"github.com/bohoslavskyi/ip-info/configs"
	"github.com/bohoslavskyi/ip-info/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	cfg *configs.Config
}

func NewHandler(cfg *configs.Config) *Handler {
	return &Handler{cfg: cfg}
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
