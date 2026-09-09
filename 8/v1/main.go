package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	start_dir := os.Args[1]
	target_filename := os.Args[2]

	wg_sub := sync.WaitGroup{}
	wg_sub.Add(1)
	go Finder(start_dir, target_filename, &wg_sub)
	wg_sub.Wait()
}

func Finder(start_dir string, target_filename string, wg_parent *sync.WaitGroup) {
	defer wg_parent.Done()

	dir_entry, err := os.ReadDir(start_dir)
	if err != nil {
		fmt.Println("Invalid Path : ", start_dir)
		return
	}

	for _, ent := range dir_entry {
		if ent.Type().IsRegular() && ent.Name() == target_filename {
			fmt.Println("Target Found in : ", start_dir)
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
