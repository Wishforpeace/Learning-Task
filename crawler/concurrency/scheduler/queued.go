package scheduler

import "crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

// 提交请求任务到requestChannel
func (s *QueueScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueueScheduler) ConfigMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

// 告诉外界有一个worker可以接受request
func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}
func (s *QueueScheduler) Run() {
	//生成channel
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		//创建请求队列和工作队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeWorker chan engine.Request
			var activeRequest engine.Request

			//当requestQ和workerQ同时有数据时
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
