package main

import "fmt"

func main() {
	var name string
	name = "hello"
	var new_name *string = &name
	fmt.Println(&name)
	fmt.Println(&new_name)

	*new_name = ""
	fmt.Println(&new_name)

	new_name = nil
	fmt.Println(&new_name)
}
