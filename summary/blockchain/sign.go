package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

// GenerateSignatures 根据私钥列表对数据进行多次签名
func GenerateSignatures(privateKeyHexs []string, data []byte) (rsvs [][]byte, err error) {
	for _, hex := range privateKeyHexs {
		rsv, err := generateSignature(hex, data)
		if err != nil {
			return nil, err
		}
		rsvs = append(rsvs, rsv)
	}
	return
}

// generateSignature 使用单个私钥对数据签名
func generateSignature(privateKeyHex string, data []byte) (rsv []byte, err error) {
	if strings.HasPrefix(privateKeyHex, "0x") {
		privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	sign, err := crypto.Sign(data, privateKey)
	if err != nil {
		return nil, err
	}
	if len(sign) != crypto.SignatureLength {
		return nil, fmt.Errorf("wrong size for signature: got %d, want %d", len(sign), crypto.SignatureLength)
	}
	return sign, nil
}

// SplitSignForRSV 对多签签名进行拆分并分类，按 R、S、V 三种类型数据分类成为三个数组
func SplitSignForRSV(rsvs [][]byte) (rs, ss [][32]byte, vs []uint8) {
	for _, rsv := range rsvs {
		var r, s [32]byte
		copy(r[:], rsv[:32])
		copy(s[:], rsv[32:64])
		v := rsv[64]
		rs = append(rs, r)
		ss = append(ss, s)
		vs = append(vs, v)
	}
	return
}

// VerifySignature 验证 metamask 等钱包的签名信息
func VerifySignature(signatureHex, message string, addr string) (bool, error) {
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return false, fmt.Errorf("decode signatureHex(%s) failed: %s", signatureHex, err)
	}
	msg := accounts.TextHash([]byte(message))
	signature[crypto.RecoveryIDOffset] -= 27
	publicKey, err := crypto.SigToPub(msg, signature)
	if err != nil {
		return false, fmt.Errorf("sigToPub failed: %s", err)
	}
	recoveredAddr := crypto.PubkeyToAddress(*publicKey)
	return CompareEVMAddresses(recoveredAddr.String(), addr), nil
}
