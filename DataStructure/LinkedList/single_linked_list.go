package main

import "fmt"

type LinkNode struct {
	Data     interface{}
	NextNode *LinkNode
}

func main() {
	node := new(LinkNode)
	node.Data = 1
	b := node
	fmt.Println(b.Data)

}
