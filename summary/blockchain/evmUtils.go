package blockchain

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// GetEVMExtraGasPricePercent 根据 chainID 获得 gasPrice 设置的增长倍数
func GetEVMExtraGasPricePercent(chainID *big.Int) decimal.Decimal {
	switch chainID.String() {
	case "56", "0x38", "97", "0x61": // BSC 链， 后两个为测试网的 chain id
		return decimal.NewFromFloat(1.1)
	case "137", "0x89", "80001", "0x13881": // Polygon 链，后两个为测试网的 chain id
		return decimal.NewFromFloat(1.5)
	default:
		return decimal.NewFromInt(1)
	}
}

// GenerateMethodID 根据函数选择器生成 methodID
func GenerateMethodID(selector string) []byte {
	transferFnSignature := []byte(selector)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	return hash.Sum(nil)[:4]
}

// IsValidAddressByRegexp validate hex address
func IsValidAddressByRegexp(iaddress interface{}) bool {
	// 利用正则表达式判断地址是否符合以太坊地址格式
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}
	zeroAddressBytes := common.FromHex(evmZeroAddress)
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV signatures R S V returned as arrays，计算签名的 R、S、V
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

// LoadPrivateKey get *ecdsa.PrivateKey and address by private key string
func LoadPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, common.Address, error) {
	if strings.HasPrefix(privateKeyStr, "0x") {
		privateKeyStr = privateKeyStr[2:]
	}
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("load PrivateKey: privateKey HexToECDSA failed: %w", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, fmt.Errorf("load PrivateKey: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKey, address, nil
}
