package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"log"
	"regexp"
	"strings"
)

//请求返回内容
type ResponseBody struct {
	Ret int    `json:"ret"`
	Act string `json:"act"`
	Msg string `json:"msg"`
}

//预约记录
type ReserveRecord struct {
	gorm.Model
	StudentID string `gorm:"unique"`
	Count     int
	RS        []ReserveInfo
}

//预约信息
type ReserveInfo struct {
	gorm.Model
	Place string
	Seat  string
	Date  string
	Start string
	End   string
}

func main() {
	reserve := ReserveRecord{
		Count: 0,
	}
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "ASP.NET_SessionId=rlf4ttuczp2lw2fo2b2omerm; _d_id=312f1451bfca3bcc9b0962a37f033f")
		fmt.Println("Visiting", r.URL.String())

	})

	go func() {
		c.OnResponse(func(resp *colly.Response) {
			//fmt.Printf("Response %s: %s\n", resp.Request.URL, resp.Body)

			info := ResponseBody{}
			err := json.Unmarshal(resp.Body, &info)
			if err != nil {
				panic(err)
			}
			//fmt.Println(info.Msg)
			reg := regexp.MustCompile(`<tbody date='(.*?)' state='1082265730' over='true'>(.*?)</tbody>`)
			if reg == nil {
				fmt.Println("reg error")
				return
			}
			result := reg.FindAllString(info.Msg, -1)
			//fmt.Println(result)

			//reserve.Count = len(result)
			//fmt.Println(reserve.Count)

			regSeat := regexp.MustCompile(`<div class='box'><a>(.*?)</a>`)
			regPlace := regexp.MustCompile(`<div class='grey'>(.*?)</div</div>`)
			regStart := regexp.MustCompile(`<div><span class='grey'>开始:</span>(.*?)</div>`)
			regEnd := regexp.MustCompile(`<div><span class='grey'>结束:</span>(.*?)</div>`)
			regDate := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
			for _, m := range result {
				//fmt.Println(result)
				rs := ReserveInfo{}

				//查找预约位置
				place := regPlace.FindAllString(m, -1)
				//fmt.Println(place)
				dom, err := goquery.NewDocumentFromReader(strings.NewReader(place[0]))
				if err != nil {
					log.Fatalln(err)
				}
				dom.Find(".grey").Each(func(i int, selection *goquery.Selection) {
					rs.Place = selection.Text()
				})

				//查找座位
				seat := regSeat.FindAllString(m, -1)
				dom, err = goquery.NewDocumentFromReader(strings.NewReader(seat[0]))
				if err != nil {
					log.Fatalln(err)
				}
				dom.Find("a").Each(func(i int, selection *goquery.Selection) {
					rs.Seat = selection.Text()
					//fmt.Println(selection.Text())
				})

				//查找预约时间
				//日期
				rs.Date = regDate.FindAllString(m, -1)[0]
				//fmt.Println(rs.Start.Date)
				start := regStart.FindAllString(m, -1)
				end := regEnd.FindAllString(m, -1)
				dom, err = goquery.NewDocumentFromReader(strings.NewReader(m))
				if err != nil {
					log.Fatalln(err)
				}
				dom.Find("date").Each(func(i int, selection *goquery.Selection) {
					fmt.Println(selection.Text())
				})
				//fmt.Println(start)
				//开始时间
				dom, err = goquery.NewDocumentFromReader(strings.NewReader(start[0]))
				if err != nil {
					log.Fatalln(err)
				}
				dom.Find(".text-primary").Each(func(i int, selection *goquery.Selection) {
					rs.Start = selection.Text()
				})
				//结束时间
				dom, err = goquery.NewDocumentFromReader(strings.NewReader(end[0]))
				if err != nil {
					log.Fatalln(err)
				}
				dom.Find(".text-primary").Each(func(i int, selection *goquery.Selection) {
					rs.End = selection.Text()
				})
				reserve.RS = append(reserve.RS, rs)
				reserve.Count++
			}

			fmt.Println(len(reserve.RS))

		})

	}()
	c.Visit("http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/center.aspx?act=get_History_resv&strat=300&StatFlag=OVER&_=1669339750019")
	fmt.Println(reserve)
}
