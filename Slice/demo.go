package main

import "fmt"

type Hello struct {
	Name string
	Age  int
}
type Info struct {
	Name string
}

func main() {
	var H1 = Hello{
		Name: "111",
		Age:  12,
	}
	var H2 = Hello{
		Name: "2222",
		Age:  32,
	}
	He := []Hello{H1, H2}
	info := make([]Info, 0)
	for i, m := range He {
		info[i].Name = m.Name
	}

	fmt.Println(info)
}
