package battles

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/darkmantle/esoc-api/http"
)

type Damage struct {
	ID    int
	Name  string
	Round int64
	Side  string
	Dmg   int
	Hits  int
	Unit  int
}

type GetDamageDataParams struct {
	ID string
}

func FetchDamageData(id string) []Damage {
	var jsonStr = http.GetJsonStringFromUrl(fmt.Sprint("https://www.esoclife.com/en/api/battle-damage/", id))
	var results []Damage

	for _, outer := range jsonStr.([]interface{}) {
		// Convert to JSON
		var jsonEncoded, jsonErr = json.Marshal(outer)
		if jsonErr != nil {
			log.Fatal(jsonErr.Error())
		}

		// Convert JSON to struct
		var battleObj Damage
		if err := json.Unmarshal(jsonEncoded, &battleObj); err != nil {
			log.Fatal(jsonErr.Error())
		}

		results = append(results, battleObj)
	}

	return results
}

func GetDamageData(params GetDamageDataParams) []Damage {
	return FetchDamageData(params.ID)
}
