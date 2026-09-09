package main

import "fmt"

func main() {
	var (
		n   int
		e   int
		arr [1000000]int
		a   int
		b   int
	)
	fmt.Scanln(&n, &e)
	for i := 0; i < e; i++ {
		fmt.Scanln(&a, &b)
		arr[a]++
		arr[b]++
	}

	for i := 1; i <= n; i++ {
		if arr[i]%2 == 1 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
