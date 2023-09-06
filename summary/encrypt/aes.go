package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/pkg/errors"
)

func AESEncrypt(data []byte, key []byte) (string, error) {
	res, err := aesEncrypt(data, key)
	if err != nil {
		return "", err
	}
	// 对 AES 加密后的数据进行 base64 加密
	str := base64.StdEncoding.EncodeToString(res)
	return str, nil
}

func AESDecrypt(data string, key []byte) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return aesDecrypt(dataByte, key)
}

// aesEncrypt 加密
func aesEncrypt(data, key []byte) ([]byte, error) {
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
	mode := cipher.NewCBCEncrypter(block, key[:blockSize])
	mode.CryptBlocks(encrypted, encryptData)
	return encrypted, nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	paddingSize := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(data, padText...)
}

// aesDecrypt 解密
func aesDecrypt(data, key []byte) ([]byte, error) {
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
