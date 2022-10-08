package main

import (
	"fmt"
	"sync"
)

//队列先进先出，和栈操作顺序相反，我们这里只实现入队，和出队操作，其他操作和栈一样。
type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

// 入队
func (queue *ArrayQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 放入切片之中，后进的元素放在数组后面
	queue.array = append(queue.array, v)

	// 队列数量+1
	queue.size += 1
}

// 出队
func (queue *ArrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.size == 0 {
		panic("empty")
	}

	// 队列最前面的元素
	v := queue.array[0]

	//// 直接原位移动，但缩容后的内存无法释放
	//for i := 0; i < queue.size; i++ {
	//	queue.array[i-1] = queue.array[i]
	//}
	//
	//// 原数组缩容
	//queue.array = queue.array[1 : queue.size-1]

	// 创建新数组
	newArray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArray[i-1] = queue.array[i]
	}

	queue.array = newArray

	queue.size -= 1
	return v
}

// 获取队列开头元素
func (queue *ArrayQueue) Peek() string {
	if queue.size == 0 {
		panic("empty")
	}
	v := queue.array[0]
	return v
}

// 获取队列大小
func (queue *ArrayQueue) Size() int {
	return queue.size
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 1
}
func main() {
	queue := new(ArrayQueue)
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
