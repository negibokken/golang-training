package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
	}
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make([]chan int64, len(roots))

	var nfiles, nbytes []int64
	for i, root := range roots {
		fileSizes[i] = make(chan int64)
		nfiles = append(nfiles, 0)
		nbytes = append(nbytes, 0)
		walkDirWrapper(root, fileSizes[i])
	}

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		var chanCount = len(roots)
		for i, root := range roots {
			select {
			case size, ok := <-fileSizes[i]:
				if !ok {
					break
				}
				nfiles[i]++
				nbytes[i] += size
			case <-tick:
				printDiskUsage(root, nfiles[i], nbytes[i])

			}

			if _, ok := <-fileSizes[i]; !ok {
				chanCount--
			}
			if chanCount == 0 {
				break loop
			}
		}
	}
	fmt.Println("--- results: ---")
	for i, root := range roots {
		printDiskUsage(root, nfiles[i], nbytes[i])
	}
}

func walkDirWrapper(root string, fileSize chan int64) {
	var n sync.WaitGroup

	n.Add(1)
	go walkDir(root, &n, fileSize)

	go func() {
		n.Wait()
		close(fileSize)
	}()
}

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}
