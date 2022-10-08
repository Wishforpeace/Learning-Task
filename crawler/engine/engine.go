package engine

import (
	"WebCrawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]
		// 抓取网页内容
		log.Printf("Fetching %s", request.Url)
		content, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("Fecher:error"+"fetching url %s:%v", request.Url, err)
			continue
		}
		// 根据任务请求中的解析函数解析网页数据
		parseResult := request.ParserFunc(content)
		// 把解析出的请求添加到请求队列
		requests = append(requests,
			parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)

		}

	}
}
