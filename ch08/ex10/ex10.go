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
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

type Data struct {
	link  []string
	depth int
}

func main() {
	depth := flag.Int("depth", 1, "crawl depth")
	flag.Parse()
	os.Args = flag.Args()

	worklist := make(chan Data)
	n := 1

	go func() {
		for _, a := range os.Args {
			d := Data{[]string{a}, 1}
			worklist <- d
		}
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth > *depth {
			continue
		}
		fmt.Println(list)
		for _, link := range list.link {

			if links.Cancelled() {
				return
			}

			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					fmt.Println("downloading", link, "contents")
					worklist <- Data{crawl(link), list.depth + 1}
				}(link)
			}
		}
	}
}
