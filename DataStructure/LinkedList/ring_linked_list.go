package main

import "fmt"

type Ring struct {
	next, prev *Ring       //  前驱和后驱节点
	Value      interface{} //  数据
}

// 初始化空的循环链表,
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 创建N个节点的循环链表
func New(n int) *Ring {
	if n < 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 0; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

// 获取第n个节点
// n为负数，向前便利，否则向后遍历
// 遍历n次，时间复杂度为O(n)
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			// 在删除时，直接指向开头
			r = r.next
		}
	}
	return r
}

// 添加节点
// 往节点A，连接一个节点，并且返回之前的节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	//fmt.Println("link", s)
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

func linkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}
	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Next()

		// 如果节点回到起点，结束
		if node == r {
			return
		}
	}
}

// 删除节点后面的n个节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// 删除
func deleteTest() {
	// 第一个节点
	r := &Ring{Value: -1}
	// 链接五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	// 打印原来的节点
	node := r

	temp := r.Unlink(3) // 删除后面两个节点
	for {

		fmt.Println("Value:", node.Value)
		node = node.Next()
		if node == r {
			break
		}
	}
	fmt.Println("---------")

	// 打印被切断的节点
	node = temp

	for {
		fmt.Println(node.Value)

		node = node.Next()

		if node == temp {
			break
		}
	}
}

// 查看循环链表长度
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.Next() {
			n++
		}
	}
	return n
}

func main() {
	linkNewTest()

	deleteTest()
}
