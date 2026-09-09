package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	strat := time.Now()
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scan := bufio.NewReader(file)

	var n int
	line, _ := scan.ReadString('\n')
	parts := strings.Fields(line)
	n, _ = strconv.Atoi(parts[0])

	arr := make([][]float64, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		line, _ := scan.ReadString('\n')
		parts := strings.Fields(line)

		for j := 0; j < n; j++ {
			arr[i][j], _ = strconv.ParseFloat(parts[j], 64)
		}
	}

	fmt.Println(Deter(arr, n))

	fmt.Println("\n\nElapsed Time : ", time.Since(strat))
}

func Deter(arr [][]float64, n int) float64 {
	var s int = 1
	for i := 0; i < n-1; i++ {
		counter, ok := FixMatrix(arr, i, n)
		if counter%2 == 1 {
			s *= -1
		}
		if !ok {
			return 0
		}
		wg := sync.WaitGroup{}
		for j := i + 1; j < n; j++ {
			var zarib float64 = arr[j][i] / arr[i][i]
			wg.Add(1)
			go Delta(arr, i, j, n, zarib, &wg)
		}
		wg.Wait()
	}

	// for i := 0; i < n; i++ {
	// 	fmt.Println(arr[i])
	// }
	var J float64 = 1

	for i := 0; i < n; i++ {
		J *= arr[i][i]
	}
	return J * float64(s)
}

func Delta(arr [][]float64, R1, R2, n int, zarib float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		arr[R2][i] += -zarib * arr[R1][i]
	}

}

func Switch(arr [][]float64, R1, R2 int) {
	row1 := arr[R1]
	row2 := arr[R2]

	arr[R1] = row2
	arr[R2] = row1

}

func FixMatrix(arr [][]float64, i, n int) (int, bool) {
	R := i
	max := math.Abs(arr[i][i])

	for r := i + 1; r < n; r++ {
		if math.Abs(arr[r][i]) > max {
			max = math.Abs(arr[r][i])
			R = r
		}
	}

	if max == 0 {
		return 0, false
	}

	if R != i {
		Switch(arr, i, R)
		return 1, true
	}

	return 0, true
}
