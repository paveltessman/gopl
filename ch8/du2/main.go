package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSizes chan<- int64) {

	for _, entry := range dirents(dir) {

		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		}

		if !info.IsDir() {
			fileSizes <- info.Size()
			continue
		}

		subdir := filepath.Join(dir, info.Name())
		walkDir(subdir, fileSizes)
	}

}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("\r%d files %.1f GB", nfiles, float64(nbytes)/1e9)
}

func main() {

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)

	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64

loop:
	for {
		select {

		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size

		case <-tick:
			printDiskUsage(nfiles, nbytes)

		}
	}

	printDiskUsage(nfiles, nbytes)
	fmt.Println()

}
