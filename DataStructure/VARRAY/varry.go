package main

import (
	"fmt"
	"sync"
)

// 可变长数组
type Array struct {
	array []int      //  固定大小的数组，用满容量和满大小的切片来代替
	len   int        // 真正长度
	cap   int        // 容量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 初始化数组
// 新建一个可变长的数组
// 利用满容量和满大小的切片来充当固定数组，结构体Array里面的字段 len  和 cap 来控制值的存取
// 不允许设置 len > cap 的可变长数组
// 时间复杂度为 O(1)
func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len lager than cap")
	}

	// 切片当数组使用
	array := make([]int, cap, cap)

	// 元数据
	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

// 增加一个元素
func (a *Array) Append(element int) {
	// 并发锁
	// 首先添加一个元素到可变长数组里，会加锁
	//  这样保证并发安全
	a.lock.Lock()
	defer a.lock.Unlock()

	// 大小等于容量，表示没有多余位置
	if a.len == a.cap {
		// 没容量，数组要扩容，扩容为原来的两倍
		// 替换原来的老数组
		newCap := 2 * a.len

		// 如果之前的容量为0，那么新容量为1
		if a.cap == 0 {
			newCap = 1
		}

		// 前者为切片长度，后者为预留长度
		newArray := make([]int, newCap, newCap)

		// 把老数组的数据移动到新数组
		// 此步骤很耗时，时间复杂度为O(n)
		// 如果容量够，时间复杂度会变为O(1)
		for k, v := range a.array {
			newArray[k] = v
		}

		// 替换数组
		a.array = newArray
		a.cap = newCap

	}

	// 把元素放到数组里
	// 将值放到数组里，
	a.array[a.len] = element

	// 真实长度+1
	// 表示真实大小又多了1
	a.len += 1
}

// 添加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		// 只需要简单遍历一下
		a.Append(v)
	}
}

// 获取指定下标元素
func (a *Array) Get(index int) int {
	// 越界
	// 只用获取下标的值，时间复杂度为O(1)
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.array[index]
}

// 获取真实长度和容量  时间复杂度为O(1)
// 返回真实长度
func (a *Array) Len() int {
	return a.len
}

// 返回容量
func (a *Array) Cap() int {
	return a.cap
}

// 辅助打印
func Print(array *Array) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 第一个元素
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, array.Get(i))
			continue
		}
		result = fmt.Sprintf("%s %d", result, array.Get(i))
	}
	result = result + "]"
	return
}

func main() {
	// 创建一个容量为3的动态数组
	a := Make(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加一个元素
	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加一个元素
	a.Append(9)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加多个元素
	a.AppendMany(8, 7)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
}
