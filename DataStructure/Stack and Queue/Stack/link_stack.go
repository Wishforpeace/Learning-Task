package main

import (
	"fmt"
	"sync"
)

// 	链表栈，后进先出
type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}

// 入栈
func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 如果栈顶为空，就增加节点
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		// 否则新元素插入链表头部
		// 原来的的链表
		// 重新赋值，方便增加
		preNode := stack.root

		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v

		// 原来的链表连接到新元素后面
		newNode.Next = preNode

		// 将新节点放到头部
		stack.root = newNode
	}
	// 栈中元素+1
	stack.size += 1
}

// 出栈
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	//	栈元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 顶部元素出栈
	topNode := stack.root
	v := topNode.Value

	// 将顶部元素的后继接上
	stack.root = topNode.Next

	// 栈中元素-1
	stack.size -= 1

	return v
}

// 获取栈顶元素
func (stack *LinkStack) Peek() string {
	if stack.size == 0 {
		panic("empty")
	}

	v := stack.root.Value
	return v
}

// 获取栈的大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// 判断栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

// 判断是否为空
func main() {
	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
	linkStack.Push("drag")
	fmt.Println("pop:", linkStack.Pop())
}
