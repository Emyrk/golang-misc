package main

import (
	"net/http"
	"net/url"
	"sync"
)

func links(u *url.URL, d *goquery.Document) (links []models.Link) {
	wg := sync.WaitGroup{}

	d.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		go func() {
			wg.Add(1)
			href, _ := item.Attr("href")
			url, _ := url.Parse(href)
			var internal bool

			if url.Host == "" {
				url.Scheme = u.Scheme
				url.Host = u.Host
			}

			links = append(links, models.Link{
				URL:       url,
				Reachable: Reachable(url.String()),
			})

			wg.Done()
		}()
	})

	wg.Wait()

	return
}

func Reachable(u string) bool {
	res, err := http.Head(u)
	if err != nil {
		return false
	}

	return res.StatusCode == 200
}
