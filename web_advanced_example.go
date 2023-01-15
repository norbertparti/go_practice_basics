package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Name string `json:"name"`
}

func parse_response_body_to_json_advanced(resp *http.Response) {
	var result Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(result.Name)
}
