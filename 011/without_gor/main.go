package main

import (
	"fmt"
	"time"
)

// تعریف ساختار prime
type prime struct {
	value   int
	counter int
}

// تابع برای بررسی اول بودن عدد
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// تولید اولین n عدد اول
func generatePrimes(n int) []int {
	primes := []int{}
	num := 2
	for len(primes) < n {
		if isPrime(num) {
			primes = append(primes, num)
		}
		num++
	}
	return primes
}

func main() {
	s := time.Now()
	// ۱۰۰ عدد اول را تولید می‌کنیم
	primes := generatePrimes(50_000)

	// آرایه‌ای از ساختارها می‌سازیم
	structs := []prime{}
	for _, p := range primes {
		structs = append(structs, prime{value: p, counter: p})
	}

	var m int
	// پردازش هر struct
	for _, s := range structs {
		pu := s.value % 16
		counter := s.counter

		// تا وقتی counter > 0 است
		for counter > 0 {
			s.value += pu
			counter--
			pu++
			pu %= 16
			if counter == 0 {
				m++
				fmt.Printf("[%d] Prime: %d | Counter: %d \n", m, s.value, counter)
				break
			}
		}
	}

	fmt.Println("\n\nElapsed : ", time.Since(s))

}
