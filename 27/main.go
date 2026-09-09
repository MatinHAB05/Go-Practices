package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// این API آخرین ۱۰ بلوک شبکه بیت‌کوین را به‌صورت JSON برمی‌گرداند
	requestUrl := "https://blockstream.info/api/v1/blocks"

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Network error:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Status Code:", res.StatusCode)
	fmt.Println(string(body))
}