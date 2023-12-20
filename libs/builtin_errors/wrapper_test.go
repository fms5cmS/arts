package builtin_errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestWrapperInterface(t *testing.T) {
	err := errors.New("path is invalid")
	qErr := NewQueryErr("file", err)
	assert.Equal(t, err, errors.Unwrap(qErr)) // 这里调用的是 errors.Unwrap()
	t.Log(qErr)                               // query file failed
}

func TestWrapperFmt(t *testing.T) {
	err := errors.New("path is invalid")
	qErr := fmt.Errorf("query file failed %w", err)
	assert.Equal(t, err, errors.Unwrap(qErr)) // 这里调用的是 errors.Unwrap()
	t.Log(qErr)                               // query file failed path is invalid
}

func TestErrChain(t *testing.T) {
	e1 := fmt.Errorf("e1: %w", io.EOF)
	e2 := fmt.Errorf("e2: %w + %w", e1, io.ErrClosedPipe) // 可以一次性包装多个 error
	e3 := fmt.Errorf("e3: %w", e2)
	e4 := fmt.Errorf("e4: %w", e3)
	fmt.Println(errors.Unwrap(e4))                               // e3: e2: e1: EOF + io: read/write on closed pipe
	fmt.Println(errors.Unwrap(errors.Unwrap(e4)))                // e2: e1: EOF + io: read/write on closed pipe
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(e4)))) // nil 注意，这里的结果是 nil！！

	fmt.Println("============================")
	e5 := fmt.Errorf("e5: %w", io.EOF)
	e6 := fmt.Errorf("e6: %w", e5)
	e7 := fmt.Errorf("e7: %w", e6)
	e8 := fmt.Errorf("e8: %w", e7)
	fmt.Println(errors.Unwrap(e8))                                                             // e7: e6: e5: EOF
	fmt.Println(errors.Unwrap(errors.Unwrap(e8)))                                              // e6: e5: EOF
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(e8))))                               // e5: EOF
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(errors.Unwrap(e8)))))                // EOF
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(errors.Unwrap(errors.Unwrap(e8)))))) // nil
}
