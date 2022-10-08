package main

import (
	"fmt"
	"sync"
)

//	数组栈，先进后出
type ArrayStack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
	lock  sync.Mutex // 并发安全使用的锁
}

//	入栈
//	将元素入栈，会先加锁实现并发安全。
//	入栈时直接把元素放在数组的最后面，然后元素数量加 1。性能损耗主要花在切片追加元素上，
//	切片如果容量不够会自动扩容，底层损耗的复杂度我们这里不计，所以时间复杂度为
//	O(1) 。
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进元素在数组后面
	stack.array = append(stack.array, v)

	// 栈元素+1
	stack.size += 1
}

//	出栈

func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	//	栈元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素
	v := stack.array[stack.size-1]

	// 切片收缩
	//	占用空间越来越大
	//stack.array = stack.array[0 : stack.size-1]

	// 创建数组
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}

	stack.array = newArray

	// 栈元素-1
	stack.size -= 1
	return v
}

// 获取栈顶元素
func (stack *ArrayStack) Peek() string {
	if stack.size == 0 {
		panic("empty")
	}
	v := stack.array[stack.size-1]
	return v
}

// 获取栈的大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

//	判断是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}
