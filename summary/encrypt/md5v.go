package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V md5 加盐，并指定加密次数
func MD5V(str string, salt string, iteration int) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先传入盐值
	h.Write(b)
	var res []byte
	res = h.Sum(nil)
	for i := 0; i < iteration-1; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	return hex.EncodeToString(res)
}
