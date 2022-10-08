package main

import (
	"fmt"
	"time"
)

type Task struct {
	F func()
}

// 创建任务
func NewTask(f func()) *Task {
	task := Task{F: f}
	return &task
}

// 任务调用
func (t *Task) RunTask() {
	t.F()
}

// 协程池
type GoroutinePool struct {
	CapNum int
	// 进任务的管道
	InChannel chan *Task
	// 任务调度的管道
	WorkChannel chan *Task
}

func NewPool(capnum int) *GoroutinePool {
	pool := GoroutinePool{
		CapNum:      capnum,
		InChannel:   make(chan *Task),
		WorkChannel: make(chan *Task),
	}
	return &pool
}

// 从InChannel管道拿到任务，放入WorkChannel
func (p *GoroutinePool) TaskInChannelOut() {
	for task := range p.InChannel {
		p.WorkChannel <- task
	}
}

// 任务执行者从WorkChannel获取任务并执行
func (p *GoroutinePool) Worker() {
	for task := range p.WorkChannel {
		task.RunTask()
	}
}

func (p *GoroutinePool) PoolRun() {
	// 任务执行
	for i := 0; i < p.CapNum; i++ {
		go p.Worker() //开启指定数量的协程执行任务
	}
	// 从InChannel管道拿到任务，本质是往WorkChannel里面添加任务
	p.TaskInChannelOut()

	close(p.WorkChannel)
	close(p.InChannel)
}

// 1.从InChannel获取任务并写入WorkChannel
// 2.从WordChannel里面获取任务并执行
// 少了往InChannel写入任务
func main() {
	cap_num := 5

	pool := NewPool(cap_num)
	go func() {
		for {
			task := NewTask(func() {
				fmt.Println(time.Now())
			})
			pool.InChannel <- task
		}
	}()

	// 任务调度
	pool.PoolRun()
}
