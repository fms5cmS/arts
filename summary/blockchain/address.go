package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"strings"
)

func convertEVMAddress2TronAddress(addr string) string {
	s := tronAddressStrPrefix + strings.TrimPrefix(addr, "0x")
	return address.HexToAddress(s).String()
	// 也可以使用下面的方式
	// hexAddrBytes := common.FromHex(addr)
	// if len(hexAddrBytes) == 0 {
	// 	return emptyAddressBase58
	// }
	// addressHex := hexAddrBytes
	// if addressHex[0] != address.TronBytePrefix {
	// 	addressHex = append([]byte{address.TronBytePrefix}, hexAddrBytes...)
	// }
	// return address.HexToAddress(hexutil.Encode(addressHex)).String()
}

func convertTronAddress2EVMAddress(addr string) string {
	addrBytes, err := address.Base58ToAddress(addr)
	if err != nil {
		return evmZeroAddress
	}
	str := strings.Replace(addrBytes.Hex(), "0x41", "0x", 1)
	return common.HexToAddress(str).Hex()
}

func CompareEVMAddresses(addr, base string) bool {
	hexAddr := common.HexToAddress(addr)
	baseAddress := common.HexToAddress(base)
	return hexAddr.Cmp(baseAddress) == 0
}
