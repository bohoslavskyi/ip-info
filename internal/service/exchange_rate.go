package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bohoslavskyi/ip-info/configs"
)

type ExchangeRateResponse struct {
	Rates map[string]float64 `json:"rates"`
}

type ExchangeRateService struct {
	cfg *configs.Config
}

func NewExchangeRateService(cfg *configs.Config) *ExchangeRateService {
	return &ExchangeRateService{cfg: cfg}
}

func (s *ExchangeRateService) GetRates(currency string) (map[string]float64, error) {
	url := fmt.Sprintf("%s/%s", s.cfg.ExchangeRateAPI, currency)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var exchangeRateResponse ExchangeRateResponse
	if err := json.Unmarshal(responseBody, &exchangeRateResponse); err != nil {
		return nil, err
	}

	return exchangeRateResponse.Rates, nil
}
