package links

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"

	"os"

	"golang.org/x/net/html"
)

var Cancel = make(chan struct{})

func Cancelled() bool {
	select {
	case <-Cancel:
		return true
	default:
		return false
	}
}

func wrappedRequest(url string, result chan *http.Response) {
	if Cancelled() {
		return
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	req = req.WithContext(ctx)

	fmt.Println("Press any key to cancel")
	go func() {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		result <- resp
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancelFunc()
		Cancel <- struct{}{}
	}()
}

// Extract links
func Extract(url string) ([]string, error) {
	if Cancelled() {
		return nil, nil
	}

	result := make(chan *http.Response)
	wrappedRequest(url, result)

	resp := <-result
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
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
				err = saveHTML(url, link.String())
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v", err)
					return
				}
				links = append(links, link.String())
			}
		}
	}

	if Cancelled() {
		return nil, nil
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if Cancelled() {
		return
	}

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
	if Cancelled() {
		return nil
	}
	fmt.Println("downloading:", link)
	result := make(chan *http.Response)
	wrappedRequest(link, result)
	resp := <-result
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
