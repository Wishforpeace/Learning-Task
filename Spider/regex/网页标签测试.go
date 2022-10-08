//package main
//
//import (
//	"fmt"
//	"regexp"
//)
//
//func main() {
//	str := `
//<!DOCTYPE html>
//<html lang="zh-CN">
//<head>
//	<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
//	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
//	<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
//	<meta charset="utf-8">
//	<link rel="shortcut icon" href="/static/img/go.ico">
//	<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
//	<meta name="author" content="polaris <polaris@studygolang.com>">
//	<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
//	<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
//</head>
//	<title></title>
//	<div>hello regexp</div>
//	<div>hello2 regexp</div>
//	<div>hello333 regexp</div>
//	<div>你欠我的钱什么时候还
//			手里没钱
//			草泥马
//	</div>
//	<body>身体</body>
//<frameset cols="15,85">
//	<frame src="/static/pkgdoc/i.html">
//	<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
//	<noframes>
//	</noframes>
//</frameset>
//</html>`
//	// 解析、编译正则表达式
//	ret := regexp.MustCompile(`<div>(.*)</div>`)
//	// (.*)不能匹配换行符
//	ret2 := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
//	// (?s)是正则表达式的模式修饰符，即Singleline(单行模式）。表示更改的含义。使它与每一个字符匹配（包括换行符\n）
//	// (.*?)是一个单元组
//	// 提取需要的信息
//
//	alls := ret.FindAllStringSubmatch(str, -1)
//	fmt.Println("alls: ", alls)
//
//	alls2 := ret2.FindAllStringSubmatch(str, -1)
//	fmt.Println("alls2 : ", alls2)
//	for _, one := range alls2 {
//		fmt.Println("one[0]=", one[0])
//		fmt.Println("one[1]=", one[1])
//	}
//}
