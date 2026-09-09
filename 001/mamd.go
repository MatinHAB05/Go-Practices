package main

import (
	"fmt"
	"myp/funct"
)

func main() {
	// var n int = 5
	// 	4
	// 6 12
	// 12 10
	// 16 6
	// -8 0
	// n := 3
	var n int64
	fmt.Scanln(&n)

	var arr []funct.Tuple = make([]funct.Tuple, n)
	// arr[0] = funct.Tuple{5, 4}
	// arr[1] = funct.Tuple{7, 10}
	// arr[2] = funct.Tuple{2, -5}
	// // arr[3] = funct.Tuple{-8, 0}
	var i int64 = 0
	for ; i < n; i++ {
		var (
			X float64
			Y float64
		)
		fmt.Scanln(&X, &Y)
		arr[i] = funct.Tuple{X, Y}
		// fmt.Println(X, Y)
	}

	center, err := funct.FindCenter(arr)
	if err != nil || !funct.IsInt(center.X) || !funct.IsInt(center.Y) {
		fmt.Println("No Answer")
		return
	}
	// fmt.Println(center)
	radius := funct.FindDistance(center, arr[0])
	var t int64 = 0
	for ; t < n; t++ {
		if funct.FindDistance(center, arr[t]) != radius {
			fmt.Println("No Answer")
			return
		}
	}
	fmt.Println(center.X, center.Y)

}
