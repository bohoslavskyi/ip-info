package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bohoslavskyi/ip-info/configs"
	"github.com/bohoslavskyi/ip-info/internal/model"
)

type IPInfoService struct {
	cfg *configs.Config
}

func NewIPInfoService(cfg *configs.Config) *IPInfoService {
	return &IPInfoService{cfg: cfg}
}

func (s *IPInfoService) GetIPInfo(ip string) (*model.IPInfo, error) {
	url := fmt.Sprintf(
		"%s/%s?fields=status,message,country,countryCode,city,lat,lon,timezone",
		s.cfg.IPInfoAPI,
		ip,
	)
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get IP info: %s", err.Error())
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err.Error())
	}

	var ipInfo model.IPInfo
	if err := json.Unmarshal(responseBody, &ipInfo); err != nil {
		return nil, fmt.Errorf("failed to convert response body to JSON: %s", err.Error())
	}

	return &ipInfo, nil
}
