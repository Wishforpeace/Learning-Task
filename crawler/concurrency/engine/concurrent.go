package engine

import "log"

// 并发引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler // 任务调度器
	WorkerCount int       //并发任务数量
}

// 任务调度器
type Scheduler interface {
	Submit(request Request)
	ConfigMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult) // worker的输出
	e.Scheduler.Run()

	// 创建goroutine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	// engine把请求任务提交给Scheduler
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	itemCount := 0
	for {
		// 接受Worker的解析结果
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item:#%d: %v\n", itemCount, item)
			itemCount++
		}
		//将Worker解析出来的Request送给Scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}

}

// 创建任务，调用worker，分发goroutine
func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
