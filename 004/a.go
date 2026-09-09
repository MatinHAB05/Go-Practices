package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	x := Tri(arr[0], arr[1], arr[2]) ||
		Tri(arr[0], arr[1], arr[3]) ||
		Tri(arr[0], arr[1], arr[4]) ||
		Tri(arr[0], arr[2], arr[3]) ||
		Tri(arr[0], arr[2], arr[4]) ||
		Tri(arr[0], arr[3], arr[4]) ||
		Tri(arr[1], arr[2], arr[3]) ||
		Tri(arr[1], arr[2], arr[4]) ||
		Tri(arr[1], arr[3], arr[4]) ||
		Tri(arr[2], arr[3], arr[4])
	if x {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func Tri(x, y, z int) bool {
	return x+y > z && x+z > y && y+z > x

}
