package unsafe_test

import (
	"testing"
	"unsafe"
)

type Programmer struct {
	name     string
	age      int
	language string
}

func TestUnsafe_ModifyPrivateMember(t *testing.T) {
	p := Programmer{name: "stefno", language: "go", age: 18}
	t.Log(p)

	// name 是结构体的第一个成员，所以可以将 &p 直接解析为 *string
	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	// Offsetof 返回结构体起始地址和指定字段起始地址之间的偏移量
	// 如果 Programmer 和代码在同一个包中，结构体的起始地址 + 字段偏移量 可以得到下一个字段的起始地址
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.age)))
	*age = 20

	// Sizeof 返回类型占用的内存，如果是切片类型，返回的是描述符的大小，而不包含其引用的内存大小！
	// 所以，如果 age、language 字段顺序换了，且先修改了 language 的值，那么，这里再修改 age 的话会有问题的！
	// 如果 Programmer 在其他包中，无法得到私有成员的变量名，可以通过 Sizeof 获得成员类型的大小，然后计算成员地址
	language := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(0) + unsafe.Sizeof(string(""))))
	*language = "Golang"

	t.Log(p)
}
