package blockchain

import "errors"

var (
	ClientInitErr = errors.New("init client failed")

	ERC20InitInstanceErr = errors.New("init ERC20 instance failed")

	/** 查询类错误，可以重试 */

	ReceiptNotFoundErr = errors.New("receipt not found")
	TxQueryErr         = errors.New("query tx failed")

	/** 交易详情类错误，不需要重试，需要记录 */

	TokenNotMatchErr  = errors.New("tx token not match")
	AmountNotMatchErr = errors.New("tx amount not match")
	FromNotMatchErr   = errors.New("tx from address not match")
	ToNotMatchErr     = errors.New("tx to address not match")
	TxFailed          = errors.New("tx failed")
)
