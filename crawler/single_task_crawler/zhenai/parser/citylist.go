package parser

import (
	"crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(bytes []byte) engine.ParseResult {
	// submatch是[][][]byte类型数据
	// 第一个[]表示匹配到多少条数据，第二个[]表示匹配的数据中要提取的内容
	//fmt.Println(string(bytes))
	submatch := cityListRe.FindAllSubmatch(bytes, -1)
	result := engine.ParseResult{}
	count := 0
	for _, item := range submatch {
		count++
		if count == 10 {
			break
		}
		//fmt.Println(string(item[1]))
		result.Items = append(result.Items, "City:"+string(item[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(item[1]),
			ParseFunc: ParseProfile,
		})
	}
	return result
}
