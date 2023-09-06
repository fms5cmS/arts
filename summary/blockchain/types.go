package blockchain

import (
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"time"
)

var (
	defaultEvmDecimals      int32 = 18
	defaultEvmPaddedLength        = 32
	defaultEvmApproveAmount       = ToWei(math.MaxInt64, defaultEvmDecimals)

	defaultLoopInterval = 2 * time.Second
	defaultLoopNum      = 30

	evmZeroAddress                = "0x0"
	evmNativeCurrencyTokenAddress = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"

	/** ERC20 转账的函数签名信息 */
	erc20TransferSelector   = "transfer(address,uint256)"
	erc20transferMethodName = "transfer"
	erc20TransferTo         = "to"
	erc20TransferAmount     = "amount"
)

type BalanceArg struct {
	Native       bool
	Decimals     int32
	UserAddress  string
	TokenAddress string
}

type LoopArg struct {
	LoopNum      int // 循环次数
	LoopInterval time.Duration
}

type ReceiptArg struct {
	TxHash string
	*LoopArg
}

type Receipt struct {
	Status int
	Gas    string
}

type CheckArg struct {
	CheckFrom   bool
	From        string
	CheckTo     bool
	To          string
	CheckAmount bool
	Amount      decimal.Decimal
	// 如果是 ERC20 转账的话，代币合约是必须检查的，所以这里不再单独列出
}

type TransferTxArg struct {
	Native       bool
	Decimals     int32
	TxHash       string
	TokenAddress string
	*LoopArg
	CheckArg
}

type Transaction struct {
	Hash   string
	From   string
	To     string
	Amount decimal.Decimal
	Gas    string
	Status int
	Data   []byte
}

type TransferArg struct {
	PrivateKey   string
	From         string
	Native       bool
	ChainID      *big.Int
	Beneficiary  string
	TokenAddress string
	Amount       *big.Int
}

type SignTransferArg struct {
	PrivateKey string
	ChainID    *big.Int
	TxBytes    []byte
}

type ContractTxArg struct {
	From            string
	ChainID         *big.Int
	ContractAddress string
	ABI             string
	MethodName      string
	Params          []any
	Value           *big.Int
}

type ApproveArg struct {
	PrivateKey   string
	ChainID      *big.Int
	TokenAddress string   // 对 PrivateKey 地址的哪种资产进行授权
	Spender      string   // 授权给谁，谁就可以对 PrivateKey 地址的指定资产进行转移
	Amount       *big.Int // 授权的额度
}

type AllowanceArg struct {
	TokenAddress string
	Decimals     int32
	Owner        string
	Spender      string
}

type TransformArg struct {
	PrivateKey   string // spender 的私钥
	From         string // owner
	Beneficiary  string // 收款地址
	TokenAddress string
	ChainID      *big.Int
	Amount       *big.Int
}
