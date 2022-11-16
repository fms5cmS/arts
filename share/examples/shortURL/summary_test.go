package shortURL

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 摘要算法生成短网址
// 每个长网址根据下面的计算方式可以得到 4 个短网址（6 位字符）,选择其中任意一个即可
func generateShortURLBySummaryDemo(url string) []string {
	one, _ := strconv.ParseInt("3FFFFFFF", 16, 64)
	result := make([]string, 0, 4)
	// 1. 长网址 md5 生成 32 位签名串
	hash := md5.New()
	hash.Write([]byte(url))
	md5Str := hex.EncodeToString(hash.Sum(nil))
	// 2. 每 8 位分割签名串，共计 4 段，每段进行处理
	length := 8
	for i := 0; i+length <= len(md5Str); i = i + length {
		subStr := md5Str[i : i+length]
		// 将 8 位字符串看作十六进制串（32 位二进制长度）
		hexValue, _ := strconv.ParseInt(subStr, 16, 64)
		// 将十六进制串与 0x3FFFFFFF（30 位二进制长度） 进行与操作，用于舍弃高两位，得到 30 位的二进制串
		andValueStr := fmt.Sprintf("%b", one&hexValue)
		// 根据每个 30 位二进制串生成一个 6 位字符串
		result = append(result, generate(andValueStr))
	}
	return result
}

func generate(str string) string {
	// 30 位的二进制串要得到 6 位字符串，分割长度为 5，将其分割为 6 段
	length := 5
	result := strings.Builder{}
	for i := len(str) - 1; i > 0; i = i - length {
		var segment string
		if i-length > 0 {
			segment = str[i-length+1 : i+1]
		} else {
			segment = str[0 : i+1]
		}
		// 将每个 5 位二进制长度的数字转为十进制
		index, _ := strconv.ParseInt(segment, 2, 8)
		// 根据十进制值从字母表中取得对应字符
		fmt.Printf("%s: %c\n", segment, chars[index])
		// 30 位二进制串的每 5 位得到一个字符，拼接起来最终得到一个 6 位字符串
		result.WriteByte(chars[index])
	}
	return result.String()
}

func TestGenerateShortURLBySummary(t *testing.T) {
	t.Log(generateShortURLBySummaryDemo("asfgjkbfdsgafjasbdgfuygadgufdg7fdtaugfabdjguiadagfsd/asdhusfhdsf.md"))
}
