package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"./links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	depth := flag.Int("depth", 1, "crawl depth")
	flag.Parse()
	fmt.Println(*depth)
	os.Args = flag.Args()
	for i, a := range os.Args {
		fmt.Println(i, a)
	}
	// worklist := make(chan []string)
	// var n int

	// n++
	// go func() { worklist <- os.Args[1:] }()

	// seen := make(map[string]bool)
	// for ; n > 0; n-- {
	// 	list := <-worklist
	// 	for _, link := range list {
	// 		if !seen[link] {
	// 			seen[link] = true
	// 			n++
	// 			go func(link string) {
	// 				worklist <- crawl(link)
	// 			}(link)
	// 		}
	// 	}
	// }
}
