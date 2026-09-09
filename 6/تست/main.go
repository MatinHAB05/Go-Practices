package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	// m := reflect.TypeOf(models.Movie{})
	// a := reflect.TypeOf(models.Actor{})
	// m_f, _ := m.FieldByName("Title")
	// tag := m_f.Tag.Get("pattern")

	// m_f2, _ := m.FieldByName("Quality")
	// tag2 := m_f2.Tag.Get("pattern")

	// a_f, _ := a.FieldByName("Name")
	// tag3 := a_f.Tag.Get("pattern")

	// fmt.Println(tag, tag2, tag3)

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	lines, _ := strconv.Atoi(parts[0])

	for i := 0; i < lines; i++ {
		Menu()
	}
}
