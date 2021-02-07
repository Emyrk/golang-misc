package main

import (
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/net/html"
)

type urlCache struct {
	urls map[string]struct{}
	sync.Mutex
}

func (v *urlCache) Set(url string) bool {
	v.Lock()
	defer v.Unlock()

	_, exist := v.urls[url]
	v.urls[url] = struct{}{}

	return !exist
}

func newURLCache() *urlCache {
	return &urlCache{
		urls: make(map[string]struct{}),
	}
}

type results struct {
	data chan string
	err  chan error
}

func newResults() *results {
	return &results{
		data: make(chan string, 1),
		err:  make(chan error, 1),
	}
}

func (r *results) close() {
	close(r.data)
	close(r.err)
}

func (r *results) WriteToSlice(s *[]string) {
	for {
		select {
		case data, open := <-r.data:
			if !open {
				return // All data done
			}
			*s = append(*s, data)
		case err := <-r.err:
			fmt.Println("e ", err)
		}
	}
}

func (r *results) Read() {
	fmt.Println("Read: start")
	counter := 0
	for c := range r.data {
		fmt.Println(c)
		counter++
	}
	fmt.Println("Read: end, counter = ", counter)
}

func crawl(url string, depth int, wg *sync.WaitGroup, cache *urlCache, res *results) {
	defer wg.Done()

	if depth == 0 || !cache.Set(url) {
		return
	}

	response, err := http.Get(url)
	if err != nil {
		res.err <- err
		return
	}
	defer response.Body.Close()

	node, err := html.Parse(response.Body)
	if err != nil {
		res.err <- err
		return
	}

	urls := grablUrls(response, node)

	res.data <- url

	for _, url := range urls {
		wg.Add(1)
		go crawl(url, depth-1, wg, cache, res)
	}
}

func grablUrls(resp *http.Response, node *html.Node) []string {
	var f func(*html.Node) []string
	var results []string

	f = func(n *html.Node) []string {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				results = append(results, link.String())
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

		return results
	}

	res := f(node)
	return res
}

// Crawl ...
func Crawl(url string, depth int) []string {
	wg := &sync.WaitGroup{}
	output := &[]string{}
	visited := newURLCache()
	results := newResults()

	go func() {
		wg.Add(1)
		go crawl(url, depth, wg, visited, results)
		wg.Wait()
		// All data is written
		close(results.data)
	}()

	results.WriteToSlice(output)
	close(results.err)
	// t := time.NewTimer()

	return *output
}

func main() {
	// r := Crawl("https://www.golang.org", 2)
	r := Crawl("www.golang.org", 2) // no schema, error should be generated and send via err

	fmt.Println(len(r))
}
