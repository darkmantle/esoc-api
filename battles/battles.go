package battles

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/darkmantle/esoc-api/http"
)

type Datetime struct {
	T time.Time
}

func (d *Datetime) UnmarshalJSON(b []byte) error {
	// 1. Unmarshal b to a Go string. json.Unmarshal uses reflection
	// to identify s as a string and then interprets b as JSON string.
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}

	// 2. Parse the result using our desired format.
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}

	// finally, assign the time value
	d.T = t

	return nil
}

type Battle struct {
	ID            int
	Type          string
	Attacker      string
	BattleField   string
	Defender      string
	Region        string
	Date          Datetime
	Val           float64
	Epic          int64
	Round         int64
	RoundAttacker int64
	RoundDefender int64
}

type GetBattleDataParams struct {
	Country string
}

func FetchBattleData() []Battle {
	var jsonStr = http.GetJsonStringFromUrl("https://www.esoclife.com/en/api/battles/1")
	var results []Battle

	// Loop through base JSON to convert
	for _, outer := range jsonStr {
		// Convert to JSON
		var jsonEncoded, jsonErr = json.Marshal(outer)
		if jsonErr != nil {
			log.Fatal(jsonErr.Error())
		}

		// Convert JSON to struct
		var battleObj Battle
		if err := json.Unmarshal(jsonEncoded, &battleObj); err != nil {
			log.Fatal(jsonErr.Error())
		}

		results = append(results, battleObj)
	}

	return results
}

func GetBattleData(params GetBattleDataParams) []Battle {
	var allResults []Battle = FetchBattleData()

	// Filter by country
	n := 0
	for _, val := range allResults {
		if strings.EqualFold(val.Attacker, params.Country) || strings.EqualFold(val.Defender, params.Country) {
			allResults[n] = val
			n++
		}
	}
	allResults = allResults[:n]

	return allResults
}
