package examples

import (
	"crypto/rand"
	"fmt"
	mrand "math/rand"
	"testing"
	"time"
)

func TestMathRand(t *testing.T) {
	// 指定 seed
	mrand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		// 生成 [0, 10) 之间的整数
		fmt.Println(mrand.Intn(10))
	}
}

// 将随机数转为了文本内容，可以转为 十六进制或 base64（最常见的用于传输 8Bit 字节码的编码方式之一）
func Key() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	// hex.EncodeToString(buf)
	// base64.StdEncoding.EncodeToString(buf)
	return fmt.Sprintf("%x", buf)
}
