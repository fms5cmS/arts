package builtin_errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"os"
	"testing"
)

/** errors.Is() */

func TestIs(t *testing.T) {
	// 生成 error 树
	e1 := fmt.Errorf("e1: %w", io.EOF)
	e2 := fmt.Errorf("e2: %w + %w", e1, io.ErrClosedPipe)
	e3 := fmt.Errorf("e3: %w", e2)
	e4 := fmt.Errorf("e4: %w", e3)

	assert.True(t, errors.Is(e4, io.EOF))
	assert.True(t, errors.Is(e4, io.ErrClosedPipe))
	assert.False(t, errors.Is(e4, io.ErrUnexpectedEOF))
}

/** errors.As() */

func TestAs(t *testing.T) {
	// 这里由于文件不存在，所以返回的 err 树中包含了 *fs.PathError
	// 所以在将 err 赋值给 *fs.PathError 类型的变量时是可以成功赋值的
	_, err := os.Open("non-existing")
	var pathErr *fs.PathError // 这里是一个 *fs.PathError 类型的变量
	assert.True(t, errors.As(err, &pathErr))
}

func TestAsInvalid(t *testing.T) {
	var origin = fmt.Errorf("error: %w", io.EOF)
	var tmp = io.ErrClosedPipe // 这里是一个 error 类型的变量
	// if errors.As(origin, &tmp) { // 这里在编译期会直接报错：second argument to errors.As should not be *error
	t.Log(tmp)
	// }
	_ = origin
}
