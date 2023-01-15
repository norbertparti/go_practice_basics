package main

import (
	"log"
	"net/http"
)

func get_web_page(url string) *http.Response {
	// Get the page
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
