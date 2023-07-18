package service

import "fmt"

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
