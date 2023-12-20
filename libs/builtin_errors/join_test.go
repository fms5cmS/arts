package builtin_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"os"
	"testing"
)

func TestJoin(t *testing.T) {
	var e1 = io.EOF
	var e2 = io.ErrClosedPipe
	var e3 = io.ErrNoProgress
	var e4 = io.ErrShortBuffer
	_, e5 := net.Dial("tcp", "invalid.address:80")
	e6 := os.Remove("/path/to/nonexistent/file")
	var e = errors.Join(e1, e2)
	e = errors.Join(e, e3)
	e = errors.Join(e, e4)
	e = errors.Join(e, e5)
	e = errors.Join(e, e6)
	t.Log(e)
	// 输出：
	// 	EOF
	//  io: read/write on closed pipe
	//  multiple Read calls return no data or error
	//  short buffer
	//  dial tcp: lookup invalid.address: no such host
	//  remove /path/to/nonexistent/file: no such file or directory

	assert.Nil(t, errors.Unwrap(e))
	assert.True(t, errors.Is(e, e6))
	assert.True(t, errors.Is(e, e5))
	assert.True(t, errors.Is(e, e4))
	assert.True(t, errors.Is(e, e1))
}
