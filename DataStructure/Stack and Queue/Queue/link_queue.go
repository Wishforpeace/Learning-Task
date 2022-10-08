package main

import (
	"fmt"
	"sync"
)

type LinkQueue struct {
	root *LinkNode // 链表起点
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

// 入队
// 将新元素存到链表末尾，每次都要遍历，时间复杂度为O(n)
func (queue *LinkQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 如果队列顶部为空，增加新的节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		// 否则插入链表的末尾
		newNode := new(LinkNode)
		newNode.Value = v

		// 遍历到链表末尾
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}

		nowNode.Next = newNode
	}

	queue.size += 1
}

// 出队
func (queue *LinkQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 	如果队列元素已空
	if queue.size == 0 {
		panic("empty")
	}
	// 顶部元素出队列
	topNode := queue.root
	v := topNode.Value

	// 将后面的链表连上
	queue.root = topNode.Next

	queue.size -= 1
	return v
}

// 弹出顶部元素
func (queue *LinkQueue) Peek() string {
	if queue.size == 0 {
		panic("empty")
	}
	v := queue.root.Value
	return v
}

func (queue *LinkQueue) Size() int {
	return queue.size
}

func (queue *LinkQueue) IsEmpty() bool {
	return queue.size == 0
}
func main() {
	queue := new(LinkQueue)
	queue.Add("cat")
	queue.Add("dog")
	queue.Add("hen")
	fmt.Println("size:", queue.Size())
	fmt.Println("pop:", queue.Remove())
	fmt.Println("pop:", queue.Remove())
	fmt.Println("size:", queue.Remove())
	queue.Add("drag")
	fmt.Println("pop:", queue.Remove())
}
