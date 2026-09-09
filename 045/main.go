package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var g singleflight.Group

// Simulates a heavy operation (e.g., database query or API call)
func getDataFromDB(key string) (string, error) {
	fmt.Printf("🔥 [DB Query] Executing real database query for key: %s\n", key)
	time.Sleep(3 * time.Second) // Simulate database delay
	fmt.Printf("🔥 [DB Query] Executing real database query for key: %s\n", key)
	fmt.Printf("🔥 [DB Query] Executing real database query for key: %s\n", key)
	fmt.Printf("🔥 [DB Query] Executing real database query for key: %s\n", key)
	fmt.Printf("🔥 [DB Query] Executing real database query for key: %s\n", key)
	time.Sleep(3 * time.Second) // Simulate database delay
	fmt.Printf("🔥 [DB Query] Executing real database query for key: %s\n", key)

	return "DB-Response-Data", nil
}

func fetchArticle(id string) (string, error, bool) {
	// g.Do ensures only ONE function call runs per unique key at a time
	val, err, shared := g.Do(id, func() (interface{}, error) {
		return getDataFromDB(id)
	})

	if err != nil {
		return "", err, false
	}

	return val.(string), nil, shared
}

func main() {
	var wg sync.WaitGroup
	requestsCount := 5

	fmt.Println("🚀 Firing 5 concurrent requests...")

	for i := 1; i <= requestsCount; i++ {
		wg.Add(1)
		go func(reqID int) {
			defer wg.Done()

			result, _, shared := fetchArticle("article-101")
			fmt.Printf("Request #%d: result = %s | shared = %t\n", reqID, result, shared)
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ All requests finished.")
}
