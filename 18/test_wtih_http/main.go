package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	// 1. Create payload struct (or map)
	// payload := map[string]interface{}{
	// 	"title":       "Inception2",
	// 	"releaseYear": 2010,
	// 	"quality":     "1080p",
	// }

	payload := map[string]interface{}{
		"cast-id":  1,
		"movie-id": 14,
	}

	// 2. Marshal to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}

	// 3. Create request with body
	req, err := http.NewRequest("POST",
		"http://localhost:8097/v0/movie-cast-link/",
		bytes.NewBuffer(body))

	// req, err := http.NewRequest("DELETE",
	// 	"http://localhost:8097/v0/movie/rem-movie",
	// 	bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("request err:", err)
		return
	}

	// Set JSON content type
	req.Header.Set("Content-Type", "application/json")

	// Dump outgoing request
	reqDump, _ := httputil.DumpRequestOut(req, true)
	fmt.Println("=== REQUEST ===")
	fmt.Println(string(reqDump))

	// 4. Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	defer res.Body.Close()

	// Dump response
	resDump, _ := httputil.DumpResponse(res, true)
	fmt.Println("=== RESPONSE ===")
	fmt.Println(string(resDump))
}
