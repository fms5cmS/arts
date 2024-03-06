package millionRequestPerMinute

import "net/http"

func HandlerRequest(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取 job
	var job Job
	JobQueue <- job
	w.WriteHeader(http.StatusOK)
}
