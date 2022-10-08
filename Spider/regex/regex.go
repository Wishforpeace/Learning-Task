package main

import (
	"regexp"
)

const text = `
my email is ccmouse@gmai.com@abc.com
email1 is abc@def.org
email2 is   kk@qq.com
email3 is  ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+\.[a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _; m := range match{}
}
