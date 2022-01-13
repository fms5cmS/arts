package examples

import (
	"fmt"
	"io"
	"net/http"
)

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

type Status struct {
	Code   int
	Reason string
}

func WriteResponse(w io.Writer, st Status, header http.Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	// 假设上面写入的时候已经发生了 error，那么下面几步调用 Write() 写入的时候最开始就会跳过
	for key, str := range header {
		fmt.Fprintf(ew, "%s: %s\r\n", key, str)
	}
	fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}
