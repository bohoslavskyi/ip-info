package handler

import (
	"net/http"

	"github.com/bohoslavskyi/ip-info/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetIPInfo(ctx *gin.Context) {
	var request model.GetIPInfoRequest
	if err := ctx.BindJSON(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusInternalServerError, map[string]string{
		"message": "Unimplemented error",
	})
}
