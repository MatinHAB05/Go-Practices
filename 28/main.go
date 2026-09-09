package main

import (
	"fmt"

	farazsms "github.com/ghaffari273/farazsms-go"
)

func main() {
	// import farazsms "github.com/ghaffari273/farazsms-go"
	sms := farazsms.New("YOUR_API_KEY")

	bal, _ := sms.Balance() // verify key — free
	fmt.Println(bal)
	sms.SendPattern("SJ3FgPrE0C", "09120000000", map[string]string{"code": "1234"}, "90008361")
	sms.SendSimple("Hello!", []string{"09120000000", "09130000000"}, "90008361")
	items, _ := sms.Inbox(1, 20)
	fmt.Println(items)

	// any of the 63 endpoints:
	sms.Request("POST", "/ws/v1/ticket", map[string]any{"department": 1, "subject": "Hi"})
}
