package main

import (
	"WebCrawler/engine"
	"WebCrawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseProfile,
	})

}
