package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var (
		a = "1"
		b int
	)
	err := json.Unmarshal([]byte(a), &b)
	fmt.Println("Unmarshal err is", err)
	fmt.Printf("Unmarshal result is %T ,%d", b, b)

	dat, err := json.Marshal(b)
	fmt.Println("Marshal err is ", err)
	fmt.Printf("Marsahl result is %s\n", string(dat))
}
