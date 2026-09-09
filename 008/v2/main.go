package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

var IsFound int64 = 0

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("USAGE : %s <start-dir> <target-file-name>\n", os.Args[0])
		return
	}
	start_dir := os.Args[1]
	target_filename := os.Args[2]

	s := time.Now()
	wg_sub := sync.WaitGroup{}
	wg_sub.Add(1)
	fmt.Printf("start directory : %s\n", start_dir)
	go Finder(start_dir, target_filename, &wg_sub)
	wg_sub.Wait()
	if IsFound == 1 {
		fmt.Println("------------------------------------------------------------------------ NOT FOUND ------------------------------------------------------------------------")
	}
	elapsed := time.Since(s)
	fmt.Println("Execution time:", elapsed)

}

func Finder(start_dir string, target_filename string, wg_parent *sync.WaitGroup) {
	defer wg_parent.Done()

	abs, err := filepath.Abs(start_dir)
	if err != nil {
		fmt.Printf("Invalid Path : %s\n", start_dir)
		return
	}

	dir_entry, err := os.ReadDir(start_dir)
	if err != nil {
		fmt.Printf("Cant Open Path : %s\n", abs)
		return
	}

	for _, ent := range dir_entry {
		if ent.Type().IsRegular() && ent.Name() == target_filename {
			fmt.Println("Target Found in : ", start_dir)
			atomic.StoreInt64(&IsFound, 1)
			return
		}
	}

	wg_sub := sync.WaitGroup{}

	for _, ent := range dir_entry {
		if ent.IsDir() {
			wg_sub.Add(1)
			go Finder(filepath.Join(start_dir, ent.Name()), target_filename, &wg_sub)
		}
	}
	wg_sub.Wait()
}
