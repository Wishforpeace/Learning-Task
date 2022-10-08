package main

import (
	"ccnu/model"
	"fmt"
	"github.com/tidwall/gjson"
	_ "github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const cookie = "ASP.NET_SessionId=yyfzdjiri33ibxjwxvlggv45; _d_id=5be8920f7250310789090c1b7f0355"

var wg sync.WaitGroup

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetInfo(count *int, ch1 chan model.User, ch2 <-chan string) {
	var user model.User
	client := &http.Client{}
	for {
		id, ok := <-ch2
		if ok {
			url := "http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/data/searchAccount.aspx?type=logonname&ReservaApply=ReservaApply&term=" + id + "&_=1647665727072"
			req, err := http.NewRequest("GET", url, nil)
			req.Header.Set("Cookie", cookie)
			resp, err := client.Do(req)
			CheckError(err)
			body, err := ioutil.ReadAll(resp.Body)
			CheckError(err)
			//defer resp.Body.Close()
			Name := gjson.Get(string(body), ".name")
			user.Name = Name.String()
			user.StudentID = id
			if user.Name != "" {
				model.DB.Create(&user)
				ch1 <- user
				*count++
			}

		}

	}

}

//// 模拟登录，获取cookie
//func Login(client *http.Client, username string, password string) {
//	var user = url.Values{}
//	//var form = struct {
//	//	Execution string,
//	//	Lt string,
//	//	EventId string,
//	//}
//	user.Set("password", password)
//	user.Set("username", username)
//	req, err := http.NewRequest("POST", "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=", strings.NewReader(user.Encode()))
//	CheckError(err)
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
//	resp, err := client.Do(req)
//	CheckError(err)
//	body, err := io.ReadAll(resp.Body)
//	CheckError(err)
//	fmt.Printf("%s", body)
//}

func StoreInSql(ch1 chan model.User) {

}
func StudentID(grades int, UserID chan string) {
	var id string
	var grade string
	grade = strconv.Itoa(grades)

	for i := 210000; i <= 219999; i++ {
		id = grade + strconv.Itoa(i)
		UserID <- id
	}
	fmt.Println("已经生成对应学号")

	wg.Done()

}
func main() {
	model.InitDB()
	runtime.GOMAXPROCS(runtime.NumCPU())
	var (
		ch1  = make(chan model.User, 10000)
		ch2  = make(chan string, 100000)
		done = make(chan interface{})
		now  = time.Now()
	)
	wg.Add(3)
	for j := 2019; j <= 2021; j++ {
		go StudentID(j, ch2)
	}
	wg.Wait()

	defer close(done)
	var count int = 0
	for i := 1; i <= 10000; i++ {
		go GetInfo(&count, ch1, ch2)
	}
	n := len(ch1)
	for i := 0; i < n; i++ {
		fmt.Printf("%s\n", <-ch1)
	}
	fmt.Printf("Took %s\n", time.Since(now))
}
