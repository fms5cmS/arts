package blockchain

import (
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"math/big"
	"time"
)

const (
	ethChain     = 1
	bscChain     = 56
	polygonChain = 137
	tronChain    = 734339968
)

const (
	defaultEvmDecimals     int32 = 18
	defaultEvmPaddedLength       = 32

	evmZeroAddress = "0x0000000000000000000000000000000000000000"
	evmEAddress    = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
)

const (
	/** ERC20 转账的函数签名信息 */
	erc20TransferSelector   = "transfer(address,uint256)"
	erc20transferMethodName = "transfer"
	erc20TransferTo         = "to"
	erc20TransferAmount     = "amount"
)

const (
	tronAddressBytePrefix = byte(0x41)
	tronAddressStrPrefix  = "41"
	tronEmptyAddressHex   = "410000000000000000000000000000000000000000"
)

var (
	unlimitedApproveAmount, _ = new(big.Int).SetString("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0)
	emptyAddressBase58        = address.HexToAddress(tronEmptyAddressHex).String()

	defaultLoopInterval = 2 * time.Second
	defaultLoopNum      = 30
)
