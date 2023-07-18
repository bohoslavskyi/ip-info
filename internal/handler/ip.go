package handler

import (
	"fmt"
	"net"
	"net/http"

	"github.com/bohoslavskyi/ip-info/internal/service"
	"github.com/gin-gonic/gin"
)

const ipAddressesLimit = 10

type GetIPInfoRequest struct {
	IPs []string `json:"ips"`
}

func (h *Handler) GetIPInfo(ctx *gin.Context) {
	var request GetIPInfoRequest
	if err := ctx.BindJSON(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := isRequestValid(request); err != nil {
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

func isRequestValid(request GetIPInfoRequest) error {
	if len(request.IPs) > ipAddressesLimit {
		return fmt.Errorf("request should contain at most %d IP addresses", ipAddressesLimit)
	}

	for _, ip := range request.IPs {
		if ipAddress := net.ParseIP(ip); ipAddress == nil {
			return fmt.Errorf("invalid IP address: %s", ip)
		}
	}

	return nil
}
