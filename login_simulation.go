package main
import(
	"fmt"
	"net/http"

	"io/ioutil"
	"strings"
)
func main(){
	//请求url
	requestUrl := "http://pass.muxi-tech.xyz/auth/api/signin"
	var(
		name string
		password string
	)
	fmt.Print("用户名")
	_,_ =fmt.Scanln(&name)
	fmt.Print("请输入密码")
	_,_ =fmt.Scanln(&password)
	data := "{\"username\":\""+name+"\",\"password\":\""+password+"\"}"
	payload := strings.NewReader(data)
	request, err := http.NewRequest("POST",requestUrl,payload)
	if err != nil{
		panic(err)
	}
		request.Header.Add("Accept","*/*")
		request.Header.Add("Accept-Encoding", "gzip, deflate")
		request.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8")
		request.Header.Add("Content-Length","54")
		request.Header.Add("Content-Type","text/plain;charset=UTF-8")
		request.Header.Add("Host", "pass.muxi-tech.xyz")
		request.Header.Add("Origin", "http://pass.muxi-tech.xyz")
		request.Header.Add("Proxy-Connection", "keep-alive")
		request.Header.Add("Referer", "http://pass.muxi-tech.xyz/intro")
		request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
		//关闭请求
		response, err := http.DefaultClient.Do(request)
		if err != nil{
			panic(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err!=nil{
			panic(err)
		}
		fmt.Println(response.Status)
		fmt.Println(string(body))
}