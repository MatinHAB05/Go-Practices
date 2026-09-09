package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"github.com/gocolly/colly"
// )

// func PageToPage(page int) string {
// 	return strconv.FormatInt(int64(page), 10)
// }
// func main() {
// 	page := 0
// 	f, err := os.OpenFile("a.txt", os.O_CREATE, 0777)
// 	if err != nil {
// 		panic(err)
// 	}
// 	os.Stdout = f
// 	// Instantiate default collector
// 	c := colly.NewCollector(
// 		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
// 		colly.AllowedDomains("herlifeapp.com"),
// 		colly.Async(true),
// 	)

// 	// On every a element which has href attribute call callback
// 	c.OnHTML(`a[href^="https://herlifeapp.com/blog/articles/"]`, func(e *colly.HTMLElement) {
// 		link := e.Attr("href")
// 		// Print link
// 		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
// 		// Visit link found on page
// 		// Only those links are visited which are in AllowedDomains
// 		page += 1
// 		c.Visit(e.Request.AbsoluteURL("https://herlifeapp.com/blog/category/menstrual-cycle" + "/" + PageToPage(page)))
// 	})

// 	// Before making a request print "Visiting ..."
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL.String())
// 	})

// 	// Start scraping on https://hackerspaces.org
// 	c.Visit("https://herlifeapp.com/blog/category/menstrual-cycle")
// }
