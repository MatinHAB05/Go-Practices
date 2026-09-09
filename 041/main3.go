package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/PuerkitoBio/goquery"
// 	"github.com/chromedp/chromedp"
// 	"github.com/gocolly/colly"
// )

// type Article struct {
// 	Title string `json:"title"`
// 	URL   string `json:"url"`
// }

// func main() {
// 	// ---------------------------------------------------------
// 	// کانفیگ باز شدن مرورگر واقعی (غیر مخفی)
// 	// ---------------------------------------------------------
// 	opts := append(chromedp.DefaultExecAllocatorOptions[:],
// 		chromedp.Flag("headless", false), // 👈 خاموش کردن حالت مخفی برای دیدن زنده مرورگر
// 		chromedp.Flag("disable-gpu", false),
// 		chromedp.WindowSize(1280, 800), // اندازه پنجره مرورگر
// 	)

// 	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
// 	defer cancel()

// 	ctx, cancel := chromedp.NewContext(allocCtx)
// 	defer cancel()

// 	ctx, cancel = context.WithTimeout(ctx, 180*time.Second)
// 	defer cancel()

// 	// ---------------------------------------------------------
// 	// Step 1: Infinite Scroll & Extract Links
// 	// ---------------------------------------------------------
// 	fmt.Println("🚀 [Step 1] Opening Chrome window and scrolling...")

// 	var renderedHTML string
// 	scrollCount := 8 // تعداد دفعات اسکرول

// 	err := chromedp.Run(ctx,
// 		chromedp.Navigate("https://www.healthywomen.org/your-health/"),
// 		chromedp.Sleep(2*time.Second),
// 	)
// 	if err != nil {
// 		log.Fatal("Navigate error:", err)
// 	}

// 	for i := 1; i <= scrollCount; i++ {
// 		fmt.Printf("   📜 Scrolling to bottom... (%d/%d)\n", i, scrollCount)
// 		err := chromedp.Run(ctx,
// 			chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight);`, nil),
// 			chromedp.Sleep(3*time.Second),
// 		)
// 		if err != nil {
// 			fmt.Println("Scroll error:", err)
// 			break
// 		}
// 	}

// 	// استخراج HTML صفحه پس از اسکرول‌ها
// 	err = chromedp.Run(ctx, chromedp.OuterHTML(`html`, &renderedHTML, chromedp.ByQuery))
// 	if err != nil {
// 		log.Fatal("HTML fetch error:", err)
// 	}

// 	// پارس کردن با GoQuery
// 	doc, err := goquery.NewDocumentFromReader(strings.NewReader(renderedHTML))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ff, err := os.OpenFile("A.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer ff.Close()

// 	visitedURLs := make(map[string]bool)
// 	var ars []string
// 	id := 0

// 	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
// 		href, exists := s.Attr("href")
// 		if !exists {
// 			return
// 		}

// 		if strings.Contains(href, "/your-wellness/") ||
// 			strings.Contains(href, "/your-health/") ||
// 			strings.Contains(href, "/your-care/") {

// 			if strings.HasPrefix(href, "/") {
// 				href = "https://www.healthywomen.org" + href
// 			}

// 			if !visitedURLs[href] {
// 				visitedURLs[href] = true
// 				id++
// 				line := strconv.Itoa(id) + " : " + href + "\n"
// 				ff.WriteString(line)
// 				ars = append(ars, href)
// 			}
// 		}
// 	})

// 	fmt.Printf("✅ [Step 1 Done] Total unique articles found: %d\n\n", len(ars))

// 	// ---------------------------------------------------------
// 	// Step 2: Scrape Metadata for Articles (with Live Output)
// 	// ---------------------------------------------------------
// 	fmt.Println("🚀 [Step 2] Scraping individual articles concurrently...")

// 	c2 := colly.NewCollector(
// 		colly.AllowedDomains("healthywomen.org", "www.healthywomen.org"),
// 		colly.Async(true),
// 		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
// 	)

// 	c2.Limit(&colly.LimitRule{
// 		DomainRegexp: `.*healthywomen\.org.*`,
// 		Parallelism:  4,
// 		Delay:        200 * time.Millisecond,
// 	})

// 	var mus sync.Mutex
// 	var atttt []Article
// 	scrapedCount := 0

// 	c2.OnHTML("html", func(h *colly.HTMLElement) {
// 		title := strings.TrimSpace(h.ChildText("h1"))
// 		url := h.Request.URL.String()

// 		if title != "" {
// 			mus.Lock()
// 			scrapedCount++
// 			atttt = append(atttt, Article{
// 				Title: title,
// 				URL:   url,
// 			})
// 			// لاگ زنده توی ترمینال
// 			fmt.Printf("   [%d/%d] Scraped: %s\n", scrapedCount, len(ars), title)
// 			mus.Unlock()
// 		}
// 	})

// 	for _, url := range ars {
// 		c2.Visit(url)
// 	}

// 	c2.Wait()

// 	// ---------------------------------------------------------
// 	// Step 3: Save Output
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
// 	fmt.Printf("\n🎉 All Done! Saved %d articles into articles.json\n", len(atttt))
// }
