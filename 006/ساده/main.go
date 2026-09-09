package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	lines, _ := strconv.Atoi(parts[0])

	for i := 0; i < lines; i++ {
		Menu()
	}
}
