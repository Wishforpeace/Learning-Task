package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "  ")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("hello %s from Go\n", *name)
	fmt.Println(fullString)
	err, result := DuplicateString(*name)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}

func DuplicateString(input string) (error, string) {
	if input == "aaa" {
		return fmt.Errorf("aaa is not allowed"), ""
	}
	return nil, input + input

}
