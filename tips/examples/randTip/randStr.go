package randTip

import (
	"golang.org/x/exp/rand"
	"strings"
	"time"
)

// RandStrSimple 根据指定的字符源 chars 生成长度为 length 的随机字符串
func RandStrSimple(chars string, length int) string {
	strBuilder := strings.Builder{}
	strBuilder.Grow(length)
	rand.Seed(uint64(time.Now().UnixNano())) // 修改随机种子
	for i := 0; i < length; i++ {
		randIdx := rand.Intn(len(chars))
		strBuilder.WriteByte(chars[randIdx])
	}
	return strBuilder.String()
}

func RandStrQuick(chars string, length, idxBits int) string {
	// idxBits = len(fmt.Sprintf("%b", len(chars)))
	strBuilder := strings.Builder{}
	strBuilder.Grow(length)
	// 	形成掩码，便于之后通过位运算得到 randIdx
	idxMask := 1<<idxBits - 1
	// 一次 rand.Uint64() 可以使用多少次
	idxNum := 63 / idxBits
	// cache 是随机位缓存
	// remainNum 是剩余随机位可使用的次数
	for i, cache, remainNum := 0, rand.Uint64(), idxNum; i < length; {
		// 剩余可用次数为 0，则重新获取随机数
		if remainNum == 0 {
			cache, remainNum = rand.Uint64(), idxNum
		}
		// 利用掩码获取有效的随机数位
		// 需要保证得到的索引值必须是小于 chars 长度的！
		// 假设 chars = "0123456789"，长度为 10，最大二进制为 1010，但是截取到的可能是 1011 即 11，这超过了 chars 的长度
		if randIdx := int(cache & uint64(idxMask)); randIdx < len(chars) {
			strBuilder.WriteByte(chars[randIdx])
			i++
		}
		// 使用下一组随机位
		cache >>= idxBits // cache = cache >> idxBits
		remainNum--
	}
	return strBuilder.String()
}
