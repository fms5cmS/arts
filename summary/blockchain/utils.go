package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"strings"
)

func IsValidAddress(chainId uint32, addr string) bool {
	if chainId == tronChain {
		_, err := address.Base58ToAddress(addr)
		return err == nil
	}
	return common.IsHexAddress(addr)
}

func NormalizeAddress(chainId uint32, addr string) string {
	addr = strings.TrimSpace(addr)
	if chainId == tronChain {
		// 兼容以太坊地址，将其转为 tron 地址
		if common.IsHexAddress(addr) {
			return convertEVMAddress2TronAddress(addr)
		}
		// tron hex address -> base58 address
		if strings.HasPrefix(addr, tronAddressStrPrefix) {
			return address.HexToAddress(addr).String()
		}
		// check tron address
		base58Addr, err := address.Base58ToAddress(addr)
		if err != nil {
			return ""
		}
		return base58Addr.String()
	}
	return common.HexToAddress(addr).Hex()
}
