package digitalSignature

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"testing"
)

// 密钥，其长度必须是 16、24、32 之一
var AESPubKey16 = []byte("ABCDABCDABCDABCD")

func TestAES(t *testing.T) {
	data := "digital signature is very hard, i can not learn very well"
	encrypted, _ := encryptByAES([]byte(data))
	afterData, _ := decryptByAES(encrypted)
	fmt.Println(string(afterData))
}

func encryptByAES(data []byte) (string, error) {
	res, err := AESEncrypt(data, AESPubKey16)
	if err != nil {
		return "", err
	}
	// 对 AES 加密后的数据进行 base64 加密
	str := base64.StdEncoding.EncodeToString(res)
	return str, nil
}

func decryptByAES(data string) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return AESDecrypt(dataByte, AESPubKey16)
}

// 加密
func AESEncrypt(data, key []byte) ([]byte, error) {
	// 根据密钥创建 block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 获取 block 的长度
	blockSize := block.BlockSize()
	// 根据 block 的长度对明文进行填充
	encryptData := pkcs7Padding(data, blockSize)
	// 初始化保存完整密文的切片
	encrypted := make([]byte, len(encryptData))
	// 使用 CBC 工作模式，同时指定 IV，这里用密钥的前 blockSize 个字符作为初始向量，解密时也必须使用相同的 IV
	// 如果是 GCM 的话，见标准库的 test 文件
	mode := cipher.NewCBCEncrypter(block, key[:blockSize])
	mode.CryptBlocks(encrypted, encryptData)
	return encrypted, nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	// 由于明文需要分组，每组的长度均为 blockSize，所以这里会先对明文的长度取余，得到最后一个 block 的长度，然后根据 blockSize 计算需要填充的长度
	// 由于 OKCS7 是冗余填充，最少填充长度为 1，最大填充长度为 blockSize
	paddingSize := blockSize - len(data)%blockSize
	// 生成填充内容，PKCS#7 在所有需要填充的位置填充完全相同的内容
	// 填充内容为需要填充的字节数，即 byte(paddingSize)
	// 假设上面计算得到需要填充的长度为 3，则需要在 data 后面填充 3 个 3
	padText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(data, padText...)
}

// 解密
func AESDecrypt(data, key []byte) ([]byte, error) {
	// 根据密钥创建 block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	mode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted := make([]byte, len(data))
	// 　解密
	mode.CryptBlocks(decrypted, data)
	// 去除填充
	return pkcs7UnPadding(decrypted)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("密文为空")
	}
	// 获得填充的长度
	unPadding := int(data[len(data)-1])
	return data[:length-unPadding], nil
}
