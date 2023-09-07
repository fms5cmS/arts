package unsafe

import (
	"unsafe"
)

// src/reflect/value.go 中 slice 和 string 的底层数据类型：

type StringHeader struct {
	Data uintptr
	Len  int
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// 实现字符串和 byte 切片的零拷贝转换
func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
