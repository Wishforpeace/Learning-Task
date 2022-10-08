package main

import "fmt"

func ArrayLink() {
	type Value struct {
		Data      string
		NextIndex int64
	}

	var array [5]Value
	array[0] = Value{"I", 3}
	array[1] = Value{"BSDSMT", 4}
	array[2] = Value{"you", 1}
	array[3] = Value{"Love", 2}
	array[4] = Value{"!", -1}

	node := array[0]
	for {
		fmt.Println(node.Data)
		if node.NextIndex == -1 {
			break
		}
		node = array[node.NextIndex]
	}
}

func main() {
	ArrayLink()
}
