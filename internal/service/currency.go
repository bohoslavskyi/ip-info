package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bohoslavskyi/ip-info/configs"
)

type Country struct {
	Currencies map[string]any `json:"currencies"`
}

type CurrencyService struct {
	cfg *configs.Config
}

func NewCurrencyService(cfg *configs.Config) *CurrencyService {
	return &CurrencyService{cfg: cfg}
}

func (s *CurrencyService) GetCurrencyByCountry(country string) (string, error) {
	url := fmt.Sprintf("%s/%s", s.cfg.CurrenciesAPI, country)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var countries []Country
	if err := json.Unmarshal(responseBody, &countries); err != nil {
		return "", err
	}

	if len(countries) == 0 || len(countries[0].Currencies) == 0 {
		return "", fmt.Errorf(
			"unable to find currency for country with name %s",
			country,
		)
	}

	currencies := []string{}
	for currency := range countries[0].Currencies {
		currencies = append(currencies, currency)
		break
	}

	return currencies[0], nil
}
