// Package main ...
package main

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Quote struct {
	Author string
	Text   string
	Tags   []string
}

func main() {
	path, exists := launcher.LookPath()
	if !exists {
		panic("WTF")
	}
	url := launcher.New().Bin(path).Headless(false).Leakless(false).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()

	page := browser.MustPage("https://quotes.toscrape.com/").MustWaitLoad()

	var allQuotes []Quote

	for {
		// 1. Wait for quotes to appear and select them
		quoteElms := page.MustElements("div.quote")

		// 2. Extract data from each quote on the current page
		// چون پوینتری هستن الان نگیری dom بعدی که روش بیاد میپرن اینا رسما
		for _, el := range quoteElms {
			text := el.MustElement("span.text").MustText()
			author := el.MustElement("small.author").MustText()

			var tags []string
			tagElms := el.MustElements("a.tag")
			for _, t := range tagElms {
				tags = append(tags, t.MustText())
			}

			allQuotes = append(allQuotes, Quote{
				Author: author,
				Text:   text,
				Tags:   tags,
			})
		}

		// 3. Check for the "Next" page button
		// quotes.toscrape.com uses li.next > a for pagination
		hasNext, nextBtn, err := page.Has("li.next > a")
		if err != nil || !hasNext {
			break // Reached the last page
		}

		// 4. Click next and wait for navigation to complete
		nextBtn.MustClick()
		// page.MustWaitLoad()

		// Small safety sleep to ensure smooth dynamic rendering
		// time.Sleep(500 * time.Millisecond)
	}

	// 5. Print results
	for i, q := range allQuotes {
		fmt.Printf("[%d] %s\n", i+1, q.Author)
		fmt.Printf("    %s\n", q.Text)
		fmt.Printf("    Tags: %s\n", strings.Join(q.Tags, " | "))
		fmt.Println(strings.Repeat("-", 40))
	}
}

// BUG
// // This example demonstrates how to use a selector to click on an element.
// func main() {
// 	path, exists := launcher.LookPath()
// 	if !exists {
// 		panic("WTF")
// 	}
// 	url := launcher.New().Bin(path).Headless(false).Leakless(false).MustLaunch()
// 	browser := rod.New().ControlURL(url).MustConnect()
// 	end := false
// 	var list []*rod.Element
// 	page := browser.MustPage("https://quotes.toscrape.com/")
// 	for i := 1; i < 50; i++ {
// 		if end {
// 			break
// 		}
// 		s := fmt.Sprintf(`a[href^="/page/%d"]`, i+1)
// 		q := "div.quote"
// 		elms := []*rod.Element(page.MustElements(q))
// 		list = append(list, elms...)
// 		ok, elm, err := page.Has(s)
// 		if err != nil || !ok {
// 			if err != nil {
// 				panic(err)
// 			}
// 			end = true
// 			continue
// 		}

// 		if err := elm.Click(proto.InputMouseButtonLeft, 1); err != nil {
// 			panic(err)
// 		}
// 	}

// 	for _, e := range list {
// 		aut := "small.author"
// 		tx := "span.text"
// 		divt := "div.tags"
// 		tags := ".tag"

// 		auth := e.MustElement(aut)
// 		text := e.MustElement(tx)
// 		Ts := e.MustElement(divt).MustElements(tags)

// 		fmt.Print(auth.MustText(), text.MustText())
// 		for _, t := range Ts {
// 			fmt.Print(t.MustText() + " | ")
// 		}
// 		fmt.Println("------------------------------")
// 	}

// 	select {}
// }

// ۱. اسکرول روی یک عنصر خاص (MustScrollIntoView)
// این کاربردی‌ترین متد است. وقتی یک عنصر را پیدا کرده‌اید و می‌خواهید مرورگر طوری اسکرول کند که آن عنصر دیده‌شود:

// Go
// // پیدا کردن عنصر و اسکرول مستقیم روی آن
// el := page.MustElement("div.quote:last-child")
// el.MustScrollIntoView()
// ۲. اسکرول با چرخ موس (Mouse.Scroll)
// اگر بخواهید مثل یک کاربر واقعی با Mouse Scroll کار کنید و مقدار مشخصی پیکسل به پایین یا بالا بروید:

// Go
// // parameters: (x, y, stepX, stepY)
// // اسکرول به اندازه ۵۰۰ پیکسل به سمت پایین
// page.Mouse.MustScroll(0, 500, 0, 5)

// // اسکرول به سمت بالا (مقادیر منفی)
// page.Mouse.MustScroll(0, -500, 0, 5)
// ۳. اسکرول با JavaScript (بسیار سریع و دقیق)
// کنترل مستقیم اسکرول صفحه با اجرای کدهای جاوااسکریپت درون مرورگر:

// Go
// // الف) اسکرول مستقیم به انتهای صفحه
// page.MustEval(`window.scrollTo(0, document.body.scrollHeight)`)

// // ب) اسکرول به مقدار مشخص (مثلاً ۵۰۰ پیکسل به پایین)
// page.MustEval(`window.scrollBy(0, 500)`)

// // ج) اسکرول نرم (Smooth Scroll)
// page.MustEval(`window.scrollTo({ top: 1000, behavior: 'smooth' })`)
// ۴. اسکرول داخل یک دایو یا کادر خاص (Scrollable Container)
// اگر داخل یک کادر کوچک (مثلاً یک Modal یا چت‌باکس که خودش Scrollbar دارد) اسکرول می‌کنید:

// Go
// // اسکرول کردن داخل یک عنصر خاص به انتهای آن کادر
// container := page.MustElement("div.chat-box")
// container.MustEval(`this.scrollTop = this.scrollHeight`)
