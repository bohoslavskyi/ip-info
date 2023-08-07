package service

import "github.com/bohoslavskyi/ip-info/internal/model"

type IPProcessor struct {
	ipInfoService       IPInfo
	currencyService     CurrencyProvider
	exchangeRateService ExchangeRateProvider
}

func NewIPProcessor(ipInfoService IPInfo, currencyService CurrencyProvider, exchangeRateService ExchangeRateProvider) *IPProcessor {
	return &IPProcessor{
		ipInfoService:       ipInfoService,
		currencyService:     currencyService,
		exchangeRateService: exchangeRateService,
	}
}

func (p *IPProcessor) Process(ip string, processedIPs chan<- model.IPDetails) {
	ipDetails := model.IPDetails{}
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

	currentTime, err := getCurrentTimeAtLocation(ipInfo.Timezone)
	if err != nil {
		ipDetails.Err = err
		processedIPs <- ipDetails
		return
	}

	processedIPs <- model.IPDetails{
		IP:          ip,
		Country:     ipInfo.Country,
		City:        ipInfo.City,
		Latitude:    ipInfo.Lat,
		Longitude:   ipInfo.Lon,
		CurrentTime: currentTime,
		Currencies:  ratesToCurrencyMap(rates, currency),
	}
}
