package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	breadthFirst(fileSystem, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func fileSystem(path string) []string {
	fmt.Println(path)

	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	var dirs []string
	for _, fi := range fileInfo {
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		if fi.IsDir() {
			dirs = append(dirs, filepath.Join(path, fi.Name()))
		}
	}
	return dirs
}
