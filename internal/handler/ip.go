package handler

import (
	"fmt"
	"net"
	"net/http"

	"github.com/bohoslavskyi/ip-info/internal/model"
	"github.com/bohoslavskyi/ip-info/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	minIPAddressesLimit = 1
	maxIPAddressesLimit = 10
)

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

	ipProcessor := service.NewIPProcessor(h.services.IPInfo, h.services.CurrencyProvider, h.services.ExchangeRateProvider)
	processedIPs := make(chan model.IPDetails)

	for _, ip := range request.IPs {
		currentIP := ip
		go ipProcessor.Process(currentIP, processedIPs)
	}

	var ipsDetails []model.IPDetails
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
	if len(request.IPs) < minIPAddressesLimit {
		return fmt.Errorf("the request should contain at least %d IP addresse(s)", minIPAddressesLimit)
	}

	if len(request.IPs) > maxIPAddressesLimit {
		return fmt.Errorf("the request should contain at most %d IP addresse(s)", maxIPAddressesLimit)
	}

	for _, ip := range request.IPs {
		if addr := net.ParseIP(ip); addr == nil {
			return fmt.Errorf("invalid IP address: %s", ip)
		}
	}

	return nil
}
