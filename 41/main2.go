package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/gocolly/colly"
// )

// type Article struct {
// 	Title string `json:"title"`
// 	URL   string `json:"url"`
// }

// func main() {
// 	// ---------------------------------------------------------
// 	// Step 1: Extract Links from Category Page
// 	// ---------------------------------------------------------
// 	ff, err := os.OpenFile("A.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer ff.Close()

// 	var mu sync.Mutex
// 	id := 0

// 	c := colly.NewCollector(
// 		colly.AllowedDomains("healthywomen.org", "www.healthywomen.org"),
// 		colly.Async(true),
// 		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
// 	)

// 	// Fixed typos and used contains (*=) for relative link support
// 	patterns := []string{
// 		`a[href*="/your-wellness/"]`,
// 		`a[href*="/your-health/"]`,
// 		`a[href*="/your-care/"]`,
// 	}

// 	visitedURLs := make(map[string]bool)
// 	var ars []string

// 	c.OnHTML(strings.Join(patterns, ","), func(h *colly.HTMLElement) {
// 		link := h.Request.AbsoluteURL(h.Attr("href"))

// 		mu.Lock()
// 		if !visitedURLs[link] && link != "" {
// 			visitedURLs[link] = true
// 			id++
// 			line := strconv.Itoa(id) + " : " + link + "\n"
// 			ff.WriteString(line)
// 			ars = append(ars, link)
// 		}
// 		mu.Unlock()
// 	})

// 	err = c.Visit("https://www.healthywomen.org/your-health/")
// 	if err != nil {
// 		fmt.Println("Visit error:", err)
// 	}

// 	c.Wait()

// 	// ---------------------------------------------------------
// 	// Step 2: Scrape Metadata for Discovered Articles
// 	// ---------------------------------------------------------
// 	c2 := colly.NewCollector(
// 		colly.AllowedDomains("healthywomen.org", "www.healthywomen.org"),
// 		colly.Async(true),
// 		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
// 	)

// 	// Rate limiter to prevent getting IP-blocked
// 	c2.Limit(&colly.LimitRule{
// 		DomainRegexp: `.*healthywomen\.org.*`,
// 		Parallelism:  4,
// 		Delay:        200 * time.Millisecond,
// 	})

// 	var mus sync.Mutex
// 	var atttt []Article

// 	c2.OnHTML("html", func(h *colly.HTMLElement) {

// 		title := strings.TrimSpace(h.ChildText("h1"))
// 		url := h.Request.URL.String()

// 		if title != "" {
// 			mus.Lock()
// 			atttt = append(atttt, Article{
// 				Title: title,
// 				URL:   url,
// 			})
// 			mus.Unlock()
// 		}
// 	})

// 	for _, url := range ars {
// 		c2.Visit(url)
// 	}

// 	c2.Wait()

// 	// ---------------------------------------------------------
// 	// Step 3: Save Output to JSON
// 	// ---------------------------------------------------------
// 	fff, err := os.OpenFile("./articles.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer fff.Close()

// 	txt, err := json.MarshalIndent(atttt, "", "  ")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fff.Write(txt)
// 	fmt.Printf("Done! Scraped %d articles.\n", len(atttt))
// }
