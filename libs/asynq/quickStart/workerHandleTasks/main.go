package main

import (
	"arts/libs/asynq/quickStart/tasks"
	"github.com/hibiken/asynq"
	"log"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// 指定并行的 worker 数量
			Concurrency: 10,
			// 指定队列的优先级
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// 其他配置见文档
		},
	)

	// 注册每个任务类型的处理函数
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDelivery, tasks.HandleEmailDeliveryTask)
	mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor())

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
