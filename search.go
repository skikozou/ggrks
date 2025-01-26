package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SearchResult struct {
	Items []struct {
		Title string
		Link  string
	}
}

func SearchOnGoogle(config *Config, words []string) (*SearchResult, error) {
	var query string
	for _, v := range words {
		query += v + "+"
	}
	query = query[:len(query)-1]

	reqURL := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s", config.GoogleSearchAPI, config.SearchEngineID, url.QueryEscape(query))

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("Http request error: %s", err.Error())
	}
	defer resp.Body.Close()

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Response decoding error: %s", err.Error())
	}

	return &result, nil
}
