//package main
//
//import (
//	"fmt"
//	"regexp"
//)
//
//func main() {
//	str1 := "abc a7c mfc cat 8ca azc cba aMc"
//	//	1.解析编译正则表达式
//	//ret := regexp.MustCompile(`a.c`) // 可以不用检查出错情况
//	//ret := regexp.MustCompile(`a[0-9]c`) 	// 中间是数字
//	//ret := regexp.MustCompile(`a\dc`) // 中间是数字
//	ret1 := regexp.MustCompile(`a[0-9A-Z]c`)
//	//
//	// 2.提取需要信息
//	alls := ret1.FindAllStringSubmatch(str1, -1)
//	fmt.Println("alls: ", alls)
//
//	str2 := "3.14 123.123 .68 hah 1.0 abc 7. ab.3 66.6 123"
//	// 解析、编译正则表达式
//	//ret2 := regexp.MustCompile(`[0-9]+\.[0-9]+`)
//	ret2 := regexp.MustCompile(`\d+\.\d+`)
//	// 提取需要的信息
//	alls2 := ret2.FindAllStringSubmatch(str2, -1)
//	fmt.Println("alls2 : ", alls2)
//}
