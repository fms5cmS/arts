package blockchain

import (
	"context"
	"github.com/shopspring/decimal"
)

type Client interface {
	// GetBalance 查询指定地址指定币种的余额
	GetBalance(ctx context.Context, arg BalanceArg) (decimal.Decimal, error)
	// GetReceipt 查询指定交易 hash 的 Receipt 信息
	GetReceipt(ctx context.Context, arg ReceiptArg) (*Receipt, error)
	// GetTransferTx 查询指定的转账交易详情
	GetTransferTx(ctx context.Context, arg TransferTxArg) (*Transaction, error)

	// Transfer 链上转账操作，Transfer = GenerateTransferData + SignTransfer + SendTransactionManually
	Transfer(ctx context.Context, arg TransferArg) (string, error)
	// GenerateTransferData 生成待签名的 转账 数据
	GenerateTransferData(ctx context.Context, arg TransferArg) ([]byte, error)
	// SignTransfer 对编码后的 转账 数据 签名
	SignTransfer(ctx context.Context, arg SignTransferArg) ([]byte, error)
	// SendTransactionManually 对待签名的转账数据进行签名并上链
	SendTransactionManually(ctx context.Context, signTxBytes []byte) (hash string, err error)
	// GenerateContractTxData 生成调用合约的数据，类似于 GenerateTransferData 方法
	GenerateContractTxData(ctx context.Context, arg ContractTxArg) ([]byte, error)

	// Approve ERC20 的授权，B 向 A 授权，以便 A 可以对 B 的某种代币进行转账操作
	Approve(ctx context.Context, arg ApproveArg) error
	// Allowance 查询剩余授权的额度，查询 A 对 B 还有多少额度可以进行操作
	Allowance(ctx context.Context, arg AllowanceArg) (decimal.Decimal, error)
	// TransferFrom A 将 B 的代币转给 C
	TransferFrom(ctx context.Context, arg TransformArg) (string, error)
}

func GetClient(chain string, rawUrl string) (Client, error) {
	switch chain {
	case "56", "0x38", "97", "0x61", // BSC，后面为测试网的 chainID
		"137", "0x89", "80001", "0x13881": // Polygon，后面为两个测试网的 chainID
		return NewEVMClient(rawUrl)
	default:
		return NewEVMClient(rawUrl)
	}
}

func initLoopArg() *LoopArg {
	return &LoopArg{
		LoopNum:      defaultLoopNum,
		LoopInterval: defaultLoopInterval,
	}
}

func checkTransferTx(tx *Transaction, checkArg CheckArg) error {
	if checkArg.CheckFrom {
		if !CompareEVMAddresses(tx.From, checkArg.From) {
			return FromNotMatchErr
		}
	}
	if checkArg.CheckTo {
		if !CompareEVMAddresses(tx.To, checkArg.To) {
			return ToNotMatchErr
		}
	}
	if checkArg.CheckAmount {
		if !tx.Amount.Equal(checkArg.Amount) {
			return AmountNotMatchErr
		}
	}
	return nil
}
