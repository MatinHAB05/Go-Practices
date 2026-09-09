package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

type Article struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	// ---------------------------------------------------------
	// Step 1: Infinite Scroll & Extract Links via Chromedp
	// ---------------------------------------------------------
	fmt.Println("Step 1: Fetching links with Infinite Scroll...")

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// مهلت کل برای اسکرول (مثلاً ۲ دقیقه)
	ctx, cancel = context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	var renderedHTML string
	scrollCount := 10 // تعداد دفعات اسکرول به انتهای صفحه (قابل تغییر)

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.healthywomen.org/your-health/"),
		chromedp.Sleep(2*time.Second),
	)
	if err != nil {
		log.Fatal("Navigate error:", err)
	}
	ش
	// حلقه اسکرول به انتهای صفحه جهت لود داینامیک مجلات
	for i := 1; i <= scrollCount; i++ {
		fmt.Printf("Scrolling down... (%d/%d)\n", i, scrollCount)
		err := chromedp.Run(ctx,
			chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight);`, nil),
			chromedp.Sleep(3*time.Second), // مهلت به JS برای رندر مطالب جدید
		)
		if err != nil {
			fmt.Println("Scroll error:", err)
			break
		}
	}

	// استخراج HTML کامل صفحه بعد از اسکرول‌ها
	err = chromedp.Run(ctx, chromedp.OuterHTML(`html`, &renderedHTML, chromedp.ByQuery))
	if err != nil {
		log.Fatal("HTML fetch error:", err)
	}

	// پارس کردن HTML حاصل از رندر مرورگر با GoQuery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(renderedHTML))
	if err != nil {
		log.Fatal(err)
	}

	ff, err := os.OpenFile("A.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer ff.Close()

	visitedURLs := make(map[string]bool)
	var ars []string
	id := 0

	// فیلتر کردن لینک‌های مجلات از روی HTML رندر شده
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		// چک کردن الگوهای مد نظر
		if strings.Contains(href, "/your-wellness/") ||
			strings.Contains(href, "/your-health/") ||
			strings.Contains(href, "/your-care/") {

			// تبدیل آدرس نسبی به کامل در صورت نیاز
			if strings.HasPrefix(href, "/") {
				href = "https://www.healthywomen.org" + href
			}

			if !visitedURLs[href] {
				visitedURLs[href] = true
				id++
				line := strconv.Itoa(id) + " : " + href + "\n"
				ff.WriteString(line)
				ars = append(ars, href)
			}
		}
	})

	fmt.Printf("Step 1 Done! Found %d unique article URLs.\n", len(ars))

	// ---------------------------------------------------------
	// Step 2: Scrape Metadata for Discovered Articles via Colly
	// ---------------------------------------------------------
	fmt.Println("Step 2: Scraping article titles...")

	c2 := colly.NewCollector(
		colly.AllowedDomains("healthywomen.org", "www.healthywomen.org"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	c2.Limit(&colly.LimitRule{
		DomainRegexp: `.*healthywomen\.org.*`,
		Parallelism:  4,
		Delay:        200 * time.Millisecond,
	})

	var mus sync.Mutex
	var atttt []Article

	c2.OnHTML("html", func(h *colly.HTMLElement) {
		title := strings.TrimSpace(h.ChildText("h1"))
		url := h.Request.URL.String()

		if title != "" {
			mus.Lock()
			atttt = append(atttt, Article{
				Title: title,
				URL:   url,
			})
			mus.Unlock()
		}
	})

	for _, url := range ars {
		c2.Visit(url)
	}

	c2.Wait()

	// ---------------------------------------------------------
	// Step 3: Save Output to JSON
	// ---------------------------------------------------------
	fff, err := os.OpenFile("./articles.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fff.Close()

	txt, err := json.MarshalIndent(atttt, "", "  ")
	if err != nil {
		panic(err)
	}

	fff.Write(txt)
	fmt.Printf("All Done! Successfully saved %d articles into articles.json.\n", len(atttt))
}
