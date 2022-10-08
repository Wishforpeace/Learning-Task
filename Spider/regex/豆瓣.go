package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// 爬取url，返回result
func HttpGetDB(url string) (result string, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err1 := client.Do(req)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	// 循环爬取整页数据
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return result, err
}
func SpiderPage(idx int) {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((idx-1)*25) + "&filter="
	// 爬取url
	result, err := HttpGetDB(url)
	if err != nil {
		fmt.Println("HttpGet err :", err)
		return
	}
	fmt.Println("result=", result)
}

func ToWork(strat, end int) {
	fmt.Printf("正在爬取%d到%d页...\n", strat, end)
	for i := strat; i < end; i++ {
		SpiderPage(i)
	}
}
func main() {
	var start, end int
	fmt.Print("请输入爬取的初始页（>=1) ：")

	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页(>=start)")
	fmt.Scan(&end)
	ToWork(start, end)
}
