package main

import "fmt"

func DeleteStr(ch string, str string) string {
	for i, _ := range ch {
		if i < len(ch)-len(str) {
			fmt.Println(ch[i : i+len(str)])
			if ch[i:i+len(str)] == str {
				ch = ch[:i] + ch[i+len(str):]
				return ch
			}
		}
	}
	return "-1"
}
func main() {
	sentence := "I love eating burger"
	str1 := "urg"
	str2 := "ae"

	sen := DeleteStr(sentence, str1)
	fmt.Println("删除", str1)
	fmt.Println(sen)
	sen = DeleteStr(sen, str2)
	fmt.Println("删除", str2)
	fmt.Println(sen)

}
