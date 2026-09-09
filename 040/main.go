package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Article struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	// Base domain restriction
	allowedDomain := "herlifeapp.com"

	// -------------------------------------------------------------
	// STEP 1: Discover Categories
	// -------------------------------------------------------------
	categoryCollector := colly.NewCollector(
		colly.AllowedDomains(allowedDomain, "www."+allowedDomain),
	)

	var categoryURLs []string

	categoryCollector.OnHTML(`a[href*="/blog/category/"]`, func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" && !contains(categoryURLs, link) {
			categoryURLs = append(categoryURLs, link)
		}
	})

	fmt.Println("🚀 Phase 1: Scraping categories...")
	categoryCollector.Visit("https://herlifeapp.com/blog/")
	fmt.Printf("✅ Found %d categories.\n\n", len(categoryURLs))
	log.Println(categoryURLs)
	fmt.Scanln()

	// -------------------------------------------------------------
	// STEP 2: Collect Article Links from Categories
	// -------------------------------------------------------------
	articleCollector := colly.NewCollector(
		colly.AllowedDomains(allowedDomain, "www."+allowedDomain),
	)

	var articleURLs []string

	articleCollector.OnHTML(`a[href*="/blog/articles/"]`, func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" && !contains(articleURLs, link) {
			articleURLs = append(articleURLs, link)
		}
	})

	fmt.Println("🚀 Phase 2: Collecting article URLs from categories...")
	for _, catURL := range categoryURLs {
		fmt.Println("🔍 Scanning category:", catURL)
		articleCollector.Visit(catURL)
	}
	fmt.Printf("✅ Found %d unique articles.\n\n", len(articleURLs))
	log.Println(articleURLs)
	fmt.Scanln()

	// -------------------------------------------------------------
	// STEP 3: Extract Content from Each Article
	// -------------------------------------------------------------
	pageCollector := colly.NewCollector(
		colly.AllowedDomains(allowedDomain, "www."+allowedDomain),
	)

	var articles []Article

	pageCollector.OnHTML("html", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.ChildText("h1"))
		if title != "" {
			articles = append(articles, Article{
				Title: title,
				URL:   e.Request.URL.String(),
			})
		}
	})

	fmt.Println("🚀 Phase 3: Extracting article content...")
	for _, artURL := range articleURLs {
		fmt.Println("📄 Scraping article:", artURL)
		pageCollector.Visit(artURL)
	}

	// Final Summary
	fmt.Printf("\n🎉 Completed! Total articles scraped: %d\n\n", len(articles))

	// -------------------------------------------------------------
	// STEP 4: Convert Output to JSON & Save
	// -------------------------------------------------------------
	jsonData, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error marshaling to JSON: %v", err)
	}

	fmt.Println("📜 JSON Output:")
	fmt.Println(string(jsonData))

	err = os.WriteFile("articles.json", jsonData, 0644)
	if err != nil {
		log.Printf("❌ Error writing JSON file: %v", err)
	} else {
		fmt.Println("\n💾 Successfully saved to articles.json")
	}

	fmt.Scanln()
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
