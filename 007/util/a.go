package util

import (
	"fmt"
	"math"
)

func PrintEven(nums []int) {
	// fmt.Println("\n\nEven:")
	var a []int

	for _, n := range nums {
		if n%2 == 0 {

			a = append(a, n)

			// fmt.Print(n, " ")
		}
	}
	fmt.Println(a)
}

func PrintOdd(nums []int) {
	// fmt.Println("\n\nOdd:")
	var a []int

	for _, n := range nums {
		if n%2 != 0 {
			// fmt.Print(n, " ")
			a = append(a, n)

		}
	}
	fmt.Println(a)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrintPrime(nums []int) {
	// fmt.Println("\n\nPrime:")
	var a []int

	for _, n := range nums {
		if isPrime(n) {
			// fmt.Print(n, " ")
			a = append(a, n)
		}
	}
	fmt.Println(a)
}

func isPerfectSquare(n int) bool {
	if n < 0 {
		return false
	}
	s := int(math.Sqrt(float64(n)))
	return s*s == n
}

func PrintPerfectSquares(nums []int) {
	var a []int
	for _, n := range nums {
		if isPerfectSquare(n) {
			// fmt.Print(n, " ")
			a = append(a, n)
		}
	}
	// fmt.Println("\n\nPerfect Squares:")
	fmt.Println(a)
}
