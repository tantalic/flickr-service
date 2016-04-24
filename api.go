package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	DEFAULT_COUNT = 20
)

type apiResponse struct {
	Version string  `json:"version"`
	Count   int     `json:"count"`
	Photos  []Photo `json:"photos"`
}

type Photo struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Square string `json:"square"`
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

func StartApiServer(c config) {
	http.HandleFunc("/", apiHandler)

	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	log.Printf("Listening on %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil || count <= 0 {
		count = DEFAULT_COUNT
	}

	photos := GetPhotos(count)

	response := apiResponse{
		Version: Version,
		Count:   len(photos),
		Photos:  photos,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-API-Version", Version)
	json.NewEncoder(w).Encode(response)
}
