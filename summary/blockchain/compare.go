package blockchain

import "github.com/ethereum/go-ethereum/common"

func CompareAddresses(addr, base string) bool {
	address := common.HexToAddress(addr)
	baseAddress := common.HexToAddress(base)
	return address.Cmp(baseAddress) == 0
}
