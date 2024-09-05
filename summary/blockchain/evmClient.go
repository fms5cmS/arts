package blockchain

import (
	"arts/summary/blockchain/contracts/erc20"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
	"time"
)

type EVMClient struct {
	rawURL string
	client *ethclient.Client
}

func NewEVMClient(rawURL string) (Client, error) {
	client, err := ethclient.Dial(rawURL)
	if err != nil {
		return nil, errors.Wrapf(err, "with %s", rawURL)
	}
	return &EVMClient{rawURL: rawURL, client: client}, nil
}

func (c *EVMClient) GetBalance(ctx context.Context, arg BalanceArg) (decimal.Decimal, error) {
	var (
		account = common.HexToAddress(arg.UserAddress)
		balance = big.NewInt(0)
		err     error
	)
	if arg.Native {
		balance, err = c.client.BalanceAt(ctx, account, nil)
	} else {
		instance, errNewInstance := erc20.NewErc20(common.HexToAddress(arg.TokenAddress), c.client)
		if errNewInstance != nil {
			err = errors.Wrap(ERC20InitInstanceErr, errNewInstance.Error())
		} else {
			balance, err = instance.BalanceOf(nil, account)
		}
	}
	return ToDecimal(balance, arg.Decimals), err
}

func (c *EVMClient) GetReceipt(ctx context.Context, arg ReceiptArg) (*Receipt, error) {
	var (
		receipt *ethtypes.Receipt
		err     error
	)
	if arg.LoopArg == nil {
		arg.LoopArg = initLoopArg()
	}
	resp := &Receipt{}
	for i := 0; i < arg.LoopNum; i++ {
		time.Sleep(arg.LoopInterval)
		receipt, err = c.client.TransactionReceipt(ctx, common.HexToHash(arg.TxHash))
		if err != nil {
			if i == arg.LoopNum-1 {
				return nil, errors.Wrapf(ReceiptNotFoundErr, "with %s", c.rawURL)
			}
			continue
		}
		if receipt.Status == 1 || receipt.Status == 0 {
			resp.Status = int(receipt.Status)
			resp.Gas = ToDecimal(CalcGasCost(receipt.GasUsed, receipt.EffectiveGasPrice), int32(defaultEvmDecimals)).String()
			break
		}
	}
	return resp, err
}

func (c *EVMClient) GetTransferTx(ctx context.Context, arg TransferTxArg) (*Transaction, error) {
	// 查询 Receipt，这里会多次轮训进行查询
	receipt, err := c.GetReceipt(ctx, ReceiptArg{TxHash: arg.TxHash, LoopArg: arg.LoopArg})
	if err != nil {
		return nil, err
	}
	if receipt.Status == 0 {
		return nil, TxFailed
	}
	var (
		contractAbi abi.ABI
		decodeData  []byte
		tx          *Transaction
	)
	// 查询基本交易信息
	tx, err = c.getBaseTransferTx(ctx, receipt, arg)
	if err != nil {
		return nil, err
	}
	if arg.Native {
		return tx, checkTransferTx(tx, arg.CheckArg)
	}
	// ERC20 的转账还需要解析 input data
	if !CompareEVMAddresses(tx.To, arg.TokenAddress) {
		return nil, TokenNotMatchErr
	}
	contractAbi, err = abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		return tx, fmt.Errorf("init ERC20 abi failed: %v", err)
	}
	inputDataStr := hex.EncodeToString(tx.Data)
	decodeData, err = hex.DecodeString(inputDataStr)
	if err != nil {
		return tx, fmt.Errorf("decode tx(%v) input data failed: %v", arg.TxHash, err)
	}
	set := make(map[string]interface{})
	if method, ok := contractAbi.Methods[erc20transferMethodName]; ok {
		if err = method.Inputs.UnpackIntoMap(set, decodeData[4:]); err != nil {
			return nil, fmt.Errorf("unpack ERC20 transfer data failed: %v", err)
		}
		if to, ok := set[erc20TransferTo].(common.Address); ok {
			tx.To = to.String()
		}
		if amount, ok := set[erc20TransferAmount].(*big.Int); ok {
			tx.Amount = ToDecimal(amount, arg.Decimals)
		}
	}
	return tx, checkTransferTx(tx, arg.CheckArg)
}

func (c *EVMClient) getBaseTransferTx(ctx context.Context, receipt *Receipt, arg TransferTxArg) (*Transaction, error) {
	var (
		tx  *ethtypes.Transaction
		msg *core.Message
		err error
	)
	tx, _, err = c.client.TransactionByHash(ctx, common.HexToHash(arg.TxHash))
	// 交易还未完成时，查到的交易接收方会是 nil，由于先查询 receipt，所以理论上这里不会出现 tx.To() == nil
	if err != nil || tx.To() == nil {
		return nil, errors.Wrap(TxQueryErr, err.Error())
	}
	msg, err = core.TransactionToMessage(tx, ethtypes.LatestSignerForChainID(tx.ChainId()), nil)
	if err != nil {
		return nil, fmt.Errorf("tx(%s) to message failed: %v", arg.TxHash, err)
	}
	return &Transaction{
		Hash:   arg.TxHash,
		From:   msg.From.String(),
		To:     tx.To().String(),
		Amount: ToDecimal(tx.Value(), arg.Decimals),
		Gas:    receipt.Gas,
		Status: receipt.Status,
		Data:   tx.Data(),
	}, err
}

func (c *EVMClient) Transfer(ctx context.Context, arg TransferArg) (string, error) {
	privateKey, err := crypto.HexToECDSA(arg.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("load private key failed: %v", err)
	}
	// 根据私钥推导公钥，再根据公钥得到以太坊地址
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	return "", fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }
	// from = crypto.PubkeyToAddress(*publicKeyECDSA)
	var signTx *ethtypes.Transaction
	tx, err := c.generateTransferTx(ctx, arg)
	if err != nil {
		return "", err
	}
	signTx, err = ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(arg.ChainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("sign tx failed: %v", err)
	}
	if err = c.client.SendTransaction(ctx, signTx); err != nil {
		return "", fmt.Errorf("send tx failed: %v, with %s", err, c.rawURL)
	}
	return signTx.Hash().String(), nil
}

func (c *EVMClient) generateTransferTx(ctx context.Context, arg TransferArg) (*ethtypes.Transaction, error) {
	var (
		value    = arg.Amount
		from     = common.HexToAddress(arg.From)
		to       = common.HexToAddress(arg.Beneficiary)
		nonce    uint64
		gasPrice *big.Int
		gasLimit uint64
		data     []byte
		err      error
	)
	if nonce, err = c.client.PendingNonceAt(ctx, from); err != nil {
		return nil, fmt.Errorf("get nonce failed: %v, with %s", err, c.rawURL)
	}
	if gasPrice, err = c.client.SuggestGasPrice(ctx); err != nil {
		return nil, fmt.Errorf("suggest gasPrice failed: %v, with %s", err, c.rawURL)
	}
	gasPrice = decimal.NewFromBigInt(gasPrice, 0).Mul(GetEVMExtraGasPricePercent(arg.ChainID)).BigInt()
	if !arg.Native {
		methodID := GenerateMethodID(erc20TransferSelector)
		paddedAddress := common.LeftPadBytes(common.HexToAddress(arg.Beneficiary).Bytes(), defaultEvmPaddedLength)
		paddedTokenAmount := common.LeftPadBytes(arg.Amount.Bytes(), defaultEvmPaddedLength)
		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedTokenAmount...)
		to = common.HexToAddress(arg.TokenAddress) // ERC20 代币转账时，to 为代币合约地址
		value = big.NewInt(0)                      // ERC20 代币转账时，value 为 0
	}
	gasLimit, err = c.client.EstimateGas(ctx, ethereum.CallMsg{To: &to, Data: data, From: from})
	if err != nil {
		// TODO add logs?
		gasLimit = uint64(210000 * 3)
	}
	tx := ethtypes.NewTx(&ethtypes.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		To:       &to,
		Gas:      gasLimit,
		Value:    value,
		Data:     data,
	})
	return tx, err
}

func (c *EVMClient) GenerateTransferData(ctx context.Context, arg TransferArg) ([]byte, error) {
	tx, err := c.generateTransferTx(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rlp.EncodeToBytes(tx)
}

func (c *EVMClient) SignTransfer(ctx context.Context, arg SignTransferArg) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(arg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("load private key failed: %v", err)
	}
	var (
		_          = ctx
		tx, signTx *ethtypes.Transaction
		signBytes  []byte
	)
	if err = rlp.DecodeBytes(arg.TxBytes, &tx); err != nil {
		return nil, fmt.Errorf("decode tx bytes failed: %v", err)
	}
	if signTx, err = ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(arg.ChainID), privateKey); err != nil {
		return nil, fmt.Errorf("sign tx failed: %v", err)
	}
	if signBytes, err = rlp.EncodeToBytes(signTx); err != nil {
		return nil, fmt.Errorf("encode tx failed: %v", err)
	}
	return signBytes, nil
}

func (c *EVMClient) SendTransactionManually(ctx context.Context, signTxBytes []byte) (hash string, err error) {
	var signTx *ethtypes.Transaction
	err = rlp.DecodeBytes(signTxBytes, &signTx)
	if err = c.client.SendTransaction(ctx, signTx); err != nil {
		return "", fmt.Errorf("send tx failed with %s: %v", c.rawURL, err)
	}
	return signTx.Hash().String(), nil
}

func (c *EVMClient) GenerateContractTxData(ctx context.Context, arg ContractTxArg) ([]byte, error) {
	var (
		nonce       uint64
		gasPrice    *big.Int
		data        []byte
		to          = common.HexToAddress(arg.ContractAddress)
		contractAbi abi.ABI
		gasLimit    uint64
		err         error
		txBytes     []byte
	)
	if nonce, err = c.client.PendingNonceAt(ctx, common.HexToAddress(arg.From)); err != nil {
		return nil, fmt.Errorf("get nonce failed: %v", err)
	}
	if gasPrice, err = c.client.SuggestGasPrice(ctx); err != nil {
		return nil, fmt.Errorf("suggest gasPrice failed: %v", err)
	}
	gasPrice = decimal.NewFromBigInt(gasPrice, 0).Mul(GetEVMExtraGasPricePercent(arg.ChainID)).BigInt()
	if contractAbi, err = abi.JSON(strings.NewReader(arg.ABI)); err != nil {
		return nil, fmt.Errorf("get abi failed: %s", err)
	}
	if data, err = contractAbi.Pack(arg.MethodName, arg.Params...); err != nil {
		return nil, fmt.Errorf("pack input data failed: %s", err)
	}
	gasLimit, err = c.client.EstimateGas(ctx, ethereum.CallMsg{From: common.HexToAddress(arg.From), Value: arg.Value, To: &to, Data: data})
	if err != nil {
		// TODO add logs?
		gasLimit = uint64(210000 * 3)
	}
	tx := ethtypes.NewTx(&ethtypes.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		To:       &to,
		Gas:      gasLimit,
		Value:    arg.Value,
		Data:     data,
	})
	if txBytes, err = rlp.EncodeToBytes(tx); err != nil {
		return nil, fmt.Errorf("tx encode to rlp bytes failed: %s", err)
	}
	return txBytes, nil
}

func (c *EVMClient) Approve(ctx context.Context, arg ApproveArg) error {
	var (
		_            = ctx
		spender      = common.HexToAddress(arg.Spender)
		tokenAddress = common.HexToAddress(arg.TokenAddress)
		instance     *erc20.Erc20
		privateKey   *ecdsa.PrivateKey
		auth         *bind.TransactOpts
		gasPrice     *big.Int
		err          error
	)
	if arg.Amount == nil {
		arg.Amount = unlimitedApproveAmount
	}
	instance, err = erc20.NewErc20(tokenAddress, c.client)
	if err != nil {
		return errors.Wrap(ERC20InitInstanceErr, err.Error())
	}
	privateKey, err = crypto.HexToECDSA(arg.PrivateKey)
	if err != nil {
		return fmt.Errorf("load private key failed: %v", err)
	}
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, arg.ChainID)
	if err != nil {
		return fmt.Errorf("new auth failed")
	}
	gasPrice, err = c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("suggest gasPrice failed: %v, with %s", err, c.rawURL)
	}
	auth.GasPrice = decimal.NewFromBigInt(gasPrice, 0).Mul(GetEVMExtraGasPricePercent(arg.ChainID)).BigInt()
	_, err = instance.Approve(auth, spender, arg.Amount)
	if err != nil {
		return fmt.Errorf("approve failed: %v", err)
	}
	return nil
}

func (c *EVMClient) Allowance(ctx context.Context, arg AllowanceArg) (decimal.Decimal, error) {
	var (
		_         = ctx
		instance  *erc20.Erc20
		allowance *big.Int
		err       error
	)
	instance, err = erc20.NewErc20(common.HexToAddress(arg.TokenAddress), c.client)
	if err != nil {
		return decimal.Zero, errors.Wrap(ERC20InitInstanceErr, err.Error())
	}
	allowance, err = instance.Allowance(nil, common.HexToAddress(arg.Owner), common.HexToAddress(arg.Spender))
	if err != nil {
		return decimal.Zero, fmt.Errorf("query allowance failed: %v", err)
	}
	return ToDecimal(allowance, arg.Decimals), nil
}

func (c *EVMClient) TransferFrom(ctx context.Context, arg TransformArg) (string, error) {
	var (
		_          = ctx
		from       = common.HexToAddress(arg.From)
		to         = common.HexToAddress(arg.Beneficiary)
		instance   *erc20.Erc20
		privateKey *ecdsa.PrivateKey
		auth       *bind.TransactOpts
		gasPrice   *big.Int
		err        error
	)
	instance, err = erc20.NewErc20(common.HexToAddress(arg.TokenAddress), c.client)
	if err != nil {
		return "", errors.Wrap(ERC20InitInstanceErr, err.Error())
	}
	privateKey, err = crypto.HexToECDSA(arg.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("load private key failed: %v", err)
	}
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, arg.ChainID)
	if err != nil {
		return "", fmt.Errorf("new auth failed")
	}
	gasPrice, err = c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("suggest gasPrice failed: %v, with %s", err, c.rawURL)
	}
	auth.GasPrice = decimal.NewFromBigInt(gasPrice, 0).Mul(GetEVMExtraGasPricePercent(arg.ChainID)).BigInt()
	tx, err := instance.TransferFrom(auth, from, to, arg.Amount)
	if err != nil {
		return "", fmt.Errorf("transfFrom failed: %v", err)
	}
	return tx.Hash().String(), nil
}
