package market

import (
	"encoding/json"
	"log"
	"sort"
	"strings"

	"github.com/darkmantle/esoc-api/http"
)

type Currency struct {
	ID       int64
	Amount   float64
	Currency string
	For      string
	Name     string
	Rate     float64
	Seller   string
}

type GetCurrencyDataParams struct {
	Currency string
	Type     string
	Limit    int64
}

func FetchCurrencyData() []Currency {
	var jsonStr = http.GetJsonStringFromUrl("https://www.esoclife.com/en/api/monetary-market/1")

	var results []Currency

	// Loop through base JSON to convert
	for _, outer := range jsonStr {
		// Convert to JSON
		var jsonEncoded, jsonErr = json.Marshal(outer)
		if jsonErr != nil {
			log.Fatal(jsonErr.Error())
		}

		// Convert JSON to struct
		var marketObj Currency
		if err := json.Unmarshal(jsonEncoded, &marketObj); err != nil {
			log.Fatal(jsonErr.Error())
		}

		results = append(results, marketObj)
	}

	return results
}

func GetCurrencyData(params GetCurrencyDataParams) []Currency {
	var allResults []Currency = FetchCurrencyData()

	if params.Limit == 0 {
		params.Limit = 10
	}

	// Filter on currency
	n := 0
	for _, val := range allResults {
		if params.Type == "To" && strings.EqualFold(val.Currency, params.Currency) {
			allResults[n] = val
			n++
		}

		if params.Type == "From" && strings.EqualFold(val.For, params.Currency) {
			allResults[n] = val
			n++
		}
	}
	allResults = allResults[:n]

	// Sort and limit
	// Sort by lowest price first
	sort.Slice(allResults, func(a, b int) bool {
		return allResults[a].Rate < allResults[b].Rate
	})

	// Return a limit
	allResults = allResults[:params.Limit]

	return allResults
}
