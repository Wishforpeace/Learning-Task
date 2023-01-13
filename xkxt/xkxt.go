package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

//func Login(c *colly.Collector) error {
//
//	err := c.Post("https://account.ccnu.edu.cn/cas/login;jsessionid=4B1D11D404464CB3DFE778D6E993403AxfPbvm?service=http://xk.ccnu.edu.cn/sso/pziotlogin",
//		map[string]string{
//			"username":  "2021214266",
//			"password":  "starsky&IMMORTAL",
//			"lt":        "LT-485771-S0nbjv21xfNfKdZIIZVuMCiFFLDRTu-account.ccnu.edu.cn",
//			"execution": "e1s1",
//			"_eventId":  "submit",
//			"submit":    "登录",
//		})
//	if err != nil {
//		return err
//	}
//	c.OnResponse(func(response *colly.Response) {
//		fmt.Println(response.StatusCode)
//		fmt.Println(response.Headers)
//
//	})
//	c.Visit("http://xk.ccnu.edu.cn/sso/pziotlogin")
//	return nil
//}
const ClassTableUrl = "http://xk.ccnu.edu.cn/jwglxt/kbcx/xskbcx_cxXsgrkb.html?gnmkdm=N2151&su=2021214266"

func main() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "")
		fmt.Println("Visiting", r.URL.String())

	})

	c.OnResponse(func(resp *colly.Response) {

	})

}
