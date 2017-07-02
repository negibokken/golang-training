package links

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"

	"os"

	"golang.org/x/net/html"
)

// Extract links
func Extract(rawurl string, origins []string) ([]string, error) {

	resp, err := http.Get(rawurl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", rawurl, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", rawurl, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				// check same origin
				flag := false
				for _, origin := range origins {
					if link.Hostname() == origin {
						flag = true
					}
				}
				if !flag {
					return
				}
				err = saveHTML(rawurl, link.String())
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v", err)
					return
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func saveHTML(origin, link string) error {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return err
	}
	u, err := url.Parse(link)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return err
	}

	dir := filepath.Join("./out", u.Host, filepath.Clean(u.Path))
	base := "[content]"

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return err
	}

	err = ioutil.WriteFile(filepath.Join(dir, base), bs, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return err
	}
	return nil
}
