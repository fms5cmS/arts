package errors

import (
	"fmt"
	"io"
	"net/http"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

func WriteResponseBefore(w io.Writer, st Status, headers []Header, body io.Reader) error {
	// 构建状态行数据，并检查错误
	if _, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason); err != nil {
		return err
	}
	// 处理 header，并每次检查错误
	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}
	// 构建结束标头部分，并检查错误
	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}
	// 将正文内容返回客户端
	_, err := io.Copy(w, body)
	return err
}

// 以上操作有很多重复性工作，可以引入一个小的包装类型来简化
// Eliminate error handling by eliminating errors！

type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	// 仅需在这里处理一遍 err
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, e.err
}

func WriteResponse(w io.Writer, st Status, header http.Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	// 底层会调用 ew 的 Write() 方法
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	// 假设上面写入的时候已经发生了 error，那么下面几步调用 Write() 写入的时候最开始就会跳过
	for key, str := range header {
		fmt.Fprintf(ew, "%s: %s\r\n", key, str)
	}
	fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}
