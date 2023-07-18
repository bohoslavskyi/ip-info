package service

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

func ratesToCurrencyMap(rates map[string]float64, currentCurrency string) []map[string]any {
	result := []map[string]any{}

	for currency, rate := range rates {
		rateTo := fmt.Sprintf("rateTo%s", currentCurrency)
		currencyMap := map[string]any{
			"currency": currency,
			rateTo:     rate,
		}
		result = append(result, currencyMap)
	}

	return result
}

func getCurrentTimeAtLocation(timezone string) (string, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", fmt.Errorf("failed to load location %s: %s", timezone, err.Error())
	}

	return time.Now().In(location).Format("02.01.2006 15:04"), nil
}
