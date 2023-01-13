package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"sync"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.bilibili.com"),
	)
	var wg sync.WaitGroup
	wg.Add(100)
	go func() {
		defer wg.Done()
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			c.Visit(e.Request.AbsoluteURL(link))
		})
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})

		c.OnResponse(func(r *colly.Response) {
			fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
		})

		c.OnError(func(r *colly.Response, err error) {
			fmt.Printf("Error %s: %v\n", r.Request.URL, err)
		})

	}()

	c.Visit("https://www.bilibili.com/")
	wg.Wait()
}
