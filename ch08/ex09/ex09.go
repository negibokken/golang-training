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

var sema = make(chan struct{}, 20)

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

	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		fileSizes[i] = make(chan int64)
		fmt.Println(i)
		go walkDir(root, &n, fileSizes[i])
	}

	go func() {
		n.Wait()
		for i := range roots {
			close(fileSizes[i])
		}
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes []int64
	for range roots {
		nfiles = append(nfiles, 0)
		nbytes = append(nbytes, 0)
	}

loop:
	for {
		// for i, root := range roots {
		root := roots[0]
		i := 0
		select {
		case size, ok := <-fileSizes[0]:
			if !ok {
				break loop
			}
			nfiles[i]++
			nbytes[i] += size
		case <-tick:
			printDiskUsage(root, nfiles[i], nbytes[i])
		}
		// }
	}
	fmt.Println("finish")
	for i, root := range roots {
		printDiskUsage(root, nfiles[i], nbytes[i])
	}
}

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}
