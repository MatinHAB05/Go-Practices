package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/gocolly/colly"
// )

// func main() {
// 	// ۱. باز کردن درست فایل برای ذخیره خروجی
// 	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	os.Stdout = f

// 	c := colly.NewCollector(
// 		colly.AllowedDomains("herlifeapp.com"),
// 		colly.Async(true),
// 	)

// 	// ۲. تنظیم محدودیت همزمانی برای جلوگیری از بلاک شدن
// 	c.Limit(&colly.LimitRule{
// 		DomainGlob:  "*herlifeapp.com*",
// 		Parallelism: 10,
// 		RandomDelay: time.Second,
// 		Delay:       time.Millisecond * 10,
// 	})

// 	// استخراج لینک مقالات
// 	c.OnHTML(`a[href^="https://herlifeapp.com/blog/articles/"]`, func(e *colly.HTMLElement) {
// 		link := e.Attr("href")
// 		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
// 	})

// 	page := 1
// 	c.OnHTML(`a.page-numbers`, func(e *colly.HTMLElement) {
// 		page++
// 		if page >= 50 {
// 			return
// 		}
// 		nextPageURL := "https://herlifeapp.com/blog/category/menstrual-cycle/page/" + strconv.Itoa(page)
// 		e.Request.Visit(nextPageURL)
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting:", r.URL.String())
// 	})

// 	c.OnError(func(r *colly.Response, err error) {
// 		fmt.Printf("Error on %s: %v\n", r.Request.URL, err)
// 	})

// 	// شروع اسکرپ
// 	c.Visit("https://herlifeapp.com/blog/category/menstrual-cycle")

// 	// ۳. ضروری برای حالت Async: منتظر ماندن تا اتمام تمام درخواست‌ها
// 	c.Wait()
// }
