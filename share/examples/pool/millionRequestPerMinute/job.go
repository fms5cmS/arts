package millionRequestPerMinute

import "fmt"

// JobQueue A buffered channel that we can send work requests on.
var JobQueue chan Job

type Job struct {
	Name string
}

// Doing 处理逻辑
func (j *Job) Doing() {
	fmt.Println(j.Name)
}
