package handler

import (
	"net/http"

	"github.com/bohoslavskyi/ip-info/internal/model"
	"github.com/bohoslavskyi/ip-info/internal/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetIPInfo(ctx *gin.Context) {
	var request model.GetIPInfoRequest
	if err := ctx.BindJSON(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ipProcessor := service.NewIPProcessor(h.cfg)
	processedIPs := make(chan service.IPDetails)

	for _, ip := range request.IPs {
		currentIP := ip
		go ipProcessor.Process(currentIP, processedIPs)
	}

	var ipsDetails []service.IPDetails
	for range request.IPs {
		processedIP := <-processedIPs
		if processedIP.Err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, processedIP.Err.Error())
			return
		}

		ipsDetails = append(ipsDetails, processedIP)
	}

	ctx.JSON(http.StatusOK, ipsDetails)
}
