package millionRequestPerMinute

type Dispatcher struct {
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{WorkerPool: make(chan chan Job, maxWorkers)}
}

func (d *Dispatcher) Run() {
	for i := 0; i < len(d.WorkerPool); i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				// 从 Worker Pool 中获取一个 Worker，并将 Job 分发给他
				// 如果没有 idle 的 Worker 的话会一直阻塞，所以分发操作需要单独开 goroutine
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}
