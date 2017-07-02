package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"net/url"

	"./links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string, origins []string) []string {
	tokens <- struct{}{}
	list, err := links.Extract(url, origins)
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

	var origins []string
	for _, rawURL := range os.Args {
		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			continue
		}
		origins = append(origins, parsedURL.Hostname())
	}
	fmt.Println(origins)

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
		for _, link := range list.link {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, origins []string) {
					links := crawl(link, origins)
					sameHosts := pickUpSameHost(links, origins)
					worklist <- Data{sameHosts, list.depth + 1}
				}(link, origins)
			}
		}
	}
}

func pickUpSameHost(links []string, origins []string) []string {
	var sameHost []string
	for _, origin := range origins {
		for _, link := range links {
			parsedURL, error := url.Parse(link)
			if error != nil {
				continue
			}
			if parsedURL.Hostname() != origin {
				continue
			}
			sameHost = append(sameHost, link)
		}
	}
	return sameHost
}
