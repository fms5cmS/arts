package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEVMAddress2TronAddress(t *testing.T) {
	testCases := []struct {
		symbol      string
		evmAddress  string
		tronAddress string
	}{
		{"USDT", "0xA614F803B6FD780986A42C78EC9C7F77E6DED13C", "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"},
		{"WBTC", "0xEFC230E125C24DE35F6290AFCAFA28D50B436536", "TXpw8XeWYeTUd4quDskoUqeQPowRh4jY65"},
		{"WETH", "0xEC51BAF14488EC651270CCC409AFDA2818AF74F2", "TXWkP3jLBqRGojUih1ShzNyDaN5Csnebok"},
		{"ZeroAddress", "0x0000000000000000000000000000000000000000", "T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.symbol, func(t *testing.T) {
			assert.Equal(t, testCase.tronAddress, convertEVMAddress2TronAddress(testCase.evmAddress))
			assert.True(t, CompareEVMAddresses(testCase.evmAddress, convertTronAddress2EVMAddress(testCase.tronAddress)))
		})
	}
}
