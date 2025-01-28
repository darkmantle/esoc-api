package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func GetJsonStringFromUrl(url string) map[string]any {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var result map[string]any

	jsonErr := json.Unmarshal([]byte(body), &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return result
}
