package concurrencyPatterns

type Conn interface {
	DoQuery(query string) Result
}

type Result int

// 多个相同操作仅取第一个操作的结果，其他操作结果会被丢弃（示例中走到 default 中）。如请求时向多个从库同时发起请求，只拿最快响应的即可。
func Query(conns []Conn, query string) Result {
	// 使用 unbuffered chan：如果在接收者准备好之前结果就返回了，发送操作就会阻塞从而走到 default 中导致最后没有任何数据返回
	// 这个 buffer 就可以保证第一个发送操作一定是成功的，而不必因为等待接收者准备好阻塞，从而进入 default 中
	ch := make(chan Result, 1)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
			default:
			}
		}(conn)
	}
	return <-ch
}
