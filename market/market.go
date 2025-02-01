package market

import (
	"encoding/json"
	"log"
	"sort"

	"github.com/darkmantle/esoc-api/http"
)

type Market struct {
	ID       int
	Name     string
	Country  string
	Currency string
	Price    float64
	Product  string
	Quality  int
	Seller   string
	Supply   int
}

type GetMarketDataParams struct {
	ProductType string
	Quality     int
	Limit       int
}

func FetchMarketData() []Market {
	var jsonStr = http.GetJsonStringFromUrl("https://www.esoclife.com/en/api/market/1")
	var results []Market

	// Loop through base JSON to convert
	for _, outer := range jsonStr {
		// Convert to JSON
		var jsonEncoded, jsonErr = json.Marshal(outer)
		if jsonErr != nil {
			log.Fatal(jsonErr.Error())
		}

		// Convert JSON to struct
		var marketObj Market
		if err := json.Unmarshal(jsonEncoded, &marketObj); err != nil {
			log.Fatal(jsonErr.Error())
		}

		results = append(results, marketObj)
	}

	return results
}

func GetMarketData(params GetMarketDataParams) []Market {
	var filtered []Market

	if params.Limit == 0 {
		params.Limit = 10
	}

	// Filter on productType and quality
	for _, val := range FetchMarketData() {
		if val.Product == params.ProductType && (val.Quality == params.Quality || params.Quality == 0) {
			filtered = append(filtered, val)
		}
	}

	// Sort by lowest price first
	sort.Slice(filtered, func(a, b int) bool {
		return filtered[a].Price < filtered[b].Price
	})

	// To avoid out of bounds error due to limit being more than length
	if len(filtered) < params.Limit {
		params.Limit = len(filtered)
	}

	// Return a limit
	filtered = filtered[:params.Limit]

	return filtered
}
