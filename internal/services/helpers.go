package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func readJsonResponse[T any](url string) (T, error) {
	var data T
	r, err := http.Get(url)
	if err != nil {
		return data, err
	}

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&data)

	return data, err
}

func debugPrint(description string, data any) {
	jsonData, _ := json.MarshalIndent(data, "", "\t")
	fmt.Printf("%s: %s", description, string(jsonData))
}
