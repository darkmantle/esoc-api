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
	var allResults []Market = FetchMarketData()

	if params.Limit == 0 {
		params.Limit = 10
	}

	// Filter on productType and quality
	n := 0
	for _, val := range allResults {
		if val.Product == params.ProductType && (val.Quality == params.Quality || params.Quality == 0) {
			allResults[n] = val
			n++
		}
	}
	allResults = allResults[:n]

	// Sort by lowest price first
	sort.Slice(allResults, func(a, b int) bool {
		return allResults[a].Price < allResults[b].Price
	})

	// Return a limit
	allResults = allResults[:params.Limit]

	return allResults
}
