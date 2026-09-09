package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func LISBinary(arr []int) int {
	var sub []int
	for _, x := range arr {
		i := sort.Search(len(sub), func(i int) bool { return sub[i] >= x })
		if i == len(sub) {
			sub = append(sub, x)
		} else {
			sub[i] = x
		}
	}
	return len(sub)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	fmt.Println(LISBinary(arr))
}
