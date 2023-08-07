package service

import (
	"github.com/bohoslavskyi/ip-info/configs"
	"github.com/bohoslavskyi/ip-info/internal/model"
)

type IPInfo interface {
	GetIPInfo(ipAddress string) (*model.IPInfo, error)
}

type CurrencyProvider interface {
	GetCurrencyByCountry(name string) (string, error)
}

type ExchangeRateProvider interface {
	GetRates(currency string) (map[string]float64, error)
}

type Service struct {
	IPInfo
	CurrencyProvider
	ExchangeRateProvider
}

func NewService(cfg *configs.Config) *Service {
	return &Service{
		IPInfo:               NewIPInfoService(cfg),
		CurrencyProvider:     NewCurrencyService(cfg),
		ExchangeRateProvider: NewExchangeRateService(cfg),
	}
}
