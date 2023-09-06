package encrypt

import "testing"

func TestMD5V(t *testing.T) {
	salt := "abcdefg"
	pwd := "12345678"
	v := MD5V(pwd, salt, 1)
	t.Logf("salt = %s, v = %s", salt, v)
	v = MD5V(pwd, salt, 2)
	t.Logf("salt = %s, v = %s", salt, v)
}
