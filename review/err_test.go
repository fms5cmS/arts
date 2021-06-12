package review

import (
	"io"
	"syscall"
)

func TestErr(t *testing.T) {
	syscall.ENAVAIL
	io.EOF
}
