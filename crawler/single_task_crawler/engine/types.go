package engine

// 请求结构
type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult //解析函数
}

// 解析结果结构
type ParseResult struct {
	Requests []Request     //解析出的请求
	Items    []interface{} //解析出的内容
}
