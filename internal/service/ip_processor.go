package service

import (
	"time"

	"github.com/bohoslavskyi/ip-info/configs"
)

type IPDetails struct {
	IP          string           `json:"ip"`
	Country     string           `json:"country"`
	City        string           `json:"city"`
	Latitude    float64          `json:"latitude"`
	Longitude   float64          `json:"longitude"`
	CurrentTime string           `json:"currentTime"`
	Currencies  []map[string]any `json:"currencies"`
	Err         error            `json:"-"`
}

type IPProcessor struct {
	ipInfoService       *IPInfoService
	currencyService     *CurrencyService
	exchangeRateService *ExchangeRateService
}

func NewIPProcessor(cfg *configs.Config) *IPProcessor {
	return &IPProcessor{
		ipInfoService:       NewIPInfoService(cfg),
		currencyService:     NewCurrencyService(cfg),
		exchangeRateService: NewExchangeRateService(cfg),
	}
}

func (p *IPProcessor) Process(ip string, processedIPs chan<- IPDetails) {
	ipDetails := IPDetails{}
	ipInfo, err := p.ipInfoService.GetIPInfo(ip)
	if err != nil {
		ipDetails.Err = err
		processedIPs <- ipDetails
		return
	}

	currency, err := p.currencyService.GetCurrencyByCountry(ipInfo.Country)
	if err != nil {
		ipDetails.Err = err
		processedIPs <- ipDetails
		return
	}

	rates, err := p.exchangeRateService.GetRates(currency)
	if err != nil {
		ipDetails.Err = err
		processedIPs <- ipDetails
		return
	}

	processedIPs <- IPDetails{
		IP:          ip,
		Country:     ipInfo.Country,
		City:        ipInfo.City,
		Latitude:    ipInfo.Lat,
		Longitude:   ipInfo.Lon,
		CurrentTime: time.Now().Format("01.12.2000 15:00"),
		Currencies:  ratesToCurrencyMap(rates, currency),
	}
}
