package main

import (
	"fmt"
	"time"
)

func main() {
	a := []string{"|\r", "/\r", "-\r", "\\\r", "|\r", "/\r", "-\r", "\\\r"}

	for {
		for _, y := range a {
			fmt.Print(y)
			time.Sleep(time.Millisecond )
		}
	}
}
