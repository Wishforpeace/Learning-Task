package main

import "fmt"

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	// 建立新节点
	node := new(LinkNode)
	node.Data = 2

	// 新节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1 //node1 链接到node的节点上

	// 新节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2

	nowNode := node

	for {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			nowNode = nowNode.NextNode
		} else {
			break
		}
	}

}
