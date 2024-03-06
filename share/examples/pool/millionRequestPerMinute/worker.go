package millionRequestPerMinute

type Worker struct {
	// 每个 Worker 会持有 WorkerPool
	WorkerPool chan chan Job
	// 该 Worker 可以处理的所有 Job
	JobChannel chan Job
	// 用于监听退出信号
	quit chan struct{}
}

func NewWorker(pool chan chan Job) Worker {
	return Worker{WorkerPool: pool, JobChannel: make(chan Job), quit: make(chan struct{})}
}

func (w *Worker) Start() {
	go func() {
		for {
			// 将当前 Worker 注册入 Worker Pool
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				job.Doing()
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- struct{}{}
	}()
}
