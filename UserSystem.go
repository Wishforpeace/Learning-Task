package main

import (
	"fmt"
	"math/rand"//生成随机数
	"net/http"
	"strconv"
	"time"
)
type User struct{
	username string
	password string
	ID string
	cookie http.Cookie
}//结构体存储用户名、密码
var mydata = make(map[string]*User)//使用map存储用户信息
func website(w http.ResponseWriter,r *http.Request ){
	for a,b :=range mydata{
		cookie, err := r.Cookie(a)
		if err == nil &&cookie.Value ==b.cookie.Value{
			fmt.Fprintf(w,"welcome!!",b.ID)
		}
	}
	w.Write([]byte("请先登录"))
}
func login(w http.ResponseWriter,r *http.Request)  {
		username :=r.FormValue("username")
		password :=r.FormValue("password")
	//确认用户是否存在
		if a, b := mydata[username];b{
			if a.password == password{//核对密码
				rand.Seed(time.Now().UnixNano())
				value := strconv.Itoa(rand.Intn(1000))
				expire := time.Now().AddDate(0, 0, 1)
				cookie := &http.Cookie{
					Name:  username,
					Value:  value,
					Expires: expire,
				}
				http.SetCookie(w,cookie)
				a.cookie = *cookie
				w.Write([]byte("登录成功"))
			}else{
				w.Write([]byte("密码错误，登录失败"))
			}
		}else{
			w.Write([]byte("用户不存在"))
		}
}
func register(w http.ResponseWriter,r *http.Request){
	var use User//结构体项目
	user := &use
	username := r.FormValue("username")
	password := r.FormValue("password")
	ID :=r.FormValue("ID")
	if username !=""&&password !=""{
		user.username = username
		user.password = password
		user.ID = ID
		mydata[username] = user//使用map，将用户身份与信息对应
		w.Write([]byte("恭喜您注册成功，请跳转至登录页面"))
	}else{
		w.Write([]byte("您未完成注册，请确保正确输入用户名及密码"))
	}
}
//修改密码
func ChangePassword(w http.ResponseWriter,r *http.Request){
	var username string
	bools :=false
	for a,b :=range mydata{
		cookie, err := r.Cookie(a)
		if err == nil && cookie.Value == b.cookie.Value{
			username = cookie.Name
			bools =true
			if bools {
				newname:=r.FormValue("newname")
				newpassword :=r.FormValue("newpassword")
				newID := r.FormValue("newID")
				user :=mydata[username]
				if newname != ""{
					user.username = newname
				}
				if newpassword != ""{
					user.password = newpassword
				}
				if newID != ""{
					user.ID = newID
				}
				w.Write([]byte("修改成功"))
			}
		}
	}
	w.Write([]byte("请登录"))
}
//查看用户信息
func WatchDetails(w http.ResponseWriter, r *http.Request){
	var username string
	bools := false
	for a,b := range mydata{
		cookie,err := r.Cookie(a)
		if(err == nil)&&(cookie.Value ==b.cookie.Value){
			username = cookie.Name
			bools = true
			if bools{
				user := mydata[username]
				fmt.Fprintf(w,"登陆成功请查看用户信息：\n用户名：%S\n昵称:%s",user.username,user.ID)
			}
		}
	}
	w.Write([]byte("请进入登录页面"))
}
func main(){
	servemux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":211",
		Handler: servemux,
	}
	servemux.HandleFunc("/website", website)
	servemux.HandleFunc("/login", login)
	servemux.HandleFunc("/register", register)
	servemux.HandleFunc("/WatchDetails", WatchDetails)
	servemux.HandleFunc("/ChangePassword", ChangePassword)
	server.ListenAndServe()
}
