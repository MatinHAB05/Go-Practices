package util

import "pp/model"

func isPrime(x int) bool {
	if x <= 1 {
		return false
	}
	if x == 2 {
		return true
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func FirstPrimes(n int) []model.Prime {
	var arr []model.Prime
	for i := 1; ; i++ {
		if isPrime(i) {
			arr = append(arr, model.Prime{Value: i, Counter: i})
		}
		if len(arr) == n {
			return arr
		}
	}
}
