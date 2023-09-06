package blockchain

import (
	"context"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
	"time"
)

type EVMClientTestSuite struct {
	suite.Suite
	evmClient Client
	ctx       context.Context

	privateKey string
	from       string

	chainID      *big.Int
	tokenAddress string
}

func (s *EVMClientTestSuite) SetupTest() {
	rawURL := "https://data-seed-prebsc-1-s1.bnbchain.org:8545"
	client, err := NewEVMClient(rawURL)
	if err != nil {
		panic(err)
	}
	s.evmClient = client
	s.ctx = context.Background()

	s.privateKey = "58c7a1f8deeced61db26b62a2fad2301a33983cee1527e3e509502f0a899ea67"
	s.from = "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89"

	s.chainID = big.NewInt(97)
	s.tokenAddress = "0x034029aFEf27c006D056Cc1eB98420aa0f3f8a1d"
}

func TestEVMClientTestSuite(t *testing.T) {
	suite.Run(t, new(EVMClientTestSuite))
}

func (s *EVMClientTestSuite) TestGetBalance() {
	testCases := []struct {
		BalanceArg
		err error
	}{
		// USDT
		{BalanceArg{Native: false, Decimals: 18, UserAddress: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89", TokenAddress: "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd"}, nil},
		// Cake
		{BalanceArg{Native: false, Decimals: 18, UserAddress: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89", TokenAddress: "0x8d008B313C1d6C7fE2982F62d32Da7507cF43551"}, nil},
		// Cake
		{BalanceArg{Native: false, Decimals: 18, UserAddress: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89", TokenAddress: "0xFa60D973F7642B748046464e165A65B7323b0DEE"}, nil},
		// tBNB
		{BalanceArg{Native: true, Decimals: 18, UserAddress: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89", TokenAddress: ""}, nil},
	}
	for i, testCase := range testCases {
		balance, err := s.evmClient.GetBalance(s.ctx, testCase.BalanceArg)
		assert.Equal(s.T(), testCase.err, errors.Cause(err))
		s.T().Logf("%d balance: %s", i, balance.String())
	}
}

func (s *EVMClientTestSuite) TestGetReceipt() {
	testCases := []struct {
		ReceiptArg
		receipt *Receipt
		err     error
	}{
		{ReceiptArg{TxHash: "0xdd09d28dbaa482d8af61bfa8c795d623713d2ae54099de4946b0cd8ed7684ee0"}, &Receipt{Status: 1, Gas: "0.000176465"}, nil},                                      // success
		{ReceiptArg{TxHash: "0xebc55e2a2d4571413c8e18c0d4eea45cabc27cb4ffbb3a28d21686b1f64e9abc"}, &Receipt{Status: 1, Gas: "0.000619762"}, nil},                                      // success
		{ReceiptArg{TxHash: "0xebc55e2a2d4571413c8e18c0d4eea45cabc27cb4ffbb3a28d21686b1f64e9ab1", LoopArg: &LoopArg{LoopNum: 1, LoopInterval: time.Second}}, nil, ReceiptNotFoundErr}, // not found
	}
	for _, testCase := range testCases {
		receipt, err := s.evmClient.GetReceipt(s.ctx, testCase.ReceiptArg)
		assert.Equal(s.T(), testCase.err, errors.Cause(err))
		if receipt != nil {
			assert.Equal(s.T(), receipt.Status, testCase.receipt.Status)
			assert.Equal(s.T(), receipt.Gas, testCase.receipt.Gas)
		}
	}
}

func (s *EVMClientTestSuite) TestGetTransferTx() {
	testCases := []struct {
		TransferTxArg
		err error
	}{
		// 完整校验
		{TransferTxArg{Native: false, Decimals: 18,
			TxHash:       "0xdd09d28dbaa482d8af61bfa8c795d623713d2ae54099de4946b0cd8ed7684ee0",
			TokenAddress: "0x034029aFEf27c006D056Cc1eB98420aa0f3f8a1d",
			CheckArg: CheckArg{
				CheckFrom: true, From: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89",
				CheckTo: true, To: "0xdf9B11c5Df465b288Af199E5291c254c81F2Ee56",
				CheckAmount: true, Amount: decimal.NewFromInt(100),
			},
		}, nil},
		{TransferTxArg{Native: false, Decimals: 18,
			TxHash:       "0x20f7b501412aa33d47dd86b2b371b2db661a52363b42104eb60d87ffb3e89cc4",
			TokenAddress: "0x034029aFEf27c006D056Cc1eB98420aa0f3f8a1d",
			CheckArg: CheckArg{
				CheckFrom: true, From: "0xdf9B11c5Df465b288Af199E5291c254c81F2Ee56",
				CheckTo: true, To: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89",
				CheckAmount: true, Amount: decimal.NewFromInt(10_000),
			},
		}, nil},
		{TransferTxArg{Native: true, Decimals: 18,
			TxHash: "0xccbae6e82a12c38d1679cfb5e86ea0b8384960c330b53e91b18c0be87eeef596", TokenAddress: "",
			CheckArg: CheckArg{
				CheckFrom: true, From: "0xaa25Aa7a19f9c426E07dee59b12f944f4d9f1DD3",
				CheckTo: true, To: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89",
				CheckAmount: true, Amount: decimal.NewFromFloat(0.1),
			},
		}, nil},
		// 四种错误的检查
		{TransferTxArg{Native: false, Decimals: 18,
			TxHash:       "0xdd09d28dbaa482d8af61bfa8c795d623713d2ae54099de4946b0cd8ed7684ee0",
			TokenAddress: "0x146E30B6D1EfdcD4E31BBBa45DDb4b233022a986",
			CheckArg: CheckArg{
				CheckFrom: true, From: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89",
				CheckTo: true, To: "0xdf9B11c5Df465b288Af199E5291c254c81F2Ee56",
				CheckAmount: true, Amount: decimal.NewFromInt(100),
			},
		}, TokenNotMatchErr},
		{TransferTxArg{Native: false, Decimals: 18,
			TxHash:       "0xdd09d28dbaa482d8af61bfa8c795d623713d2ae54099de4946b0cd8ed7684ee0",
			TokenAddress: "0x034029aFEf27c006D056Cc1eB98420aa0f3f8a1d",
			CheckArg: CheckArg{
				CheckFrom: true, From: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b88",
				CheckTo: true, To: "0xdf9B11c5Df465b288Af199E5291c254c81F2Ee56",
				CheckAmount: true, Amount: decimal.NewFromInt(100),
			},
		}, FromNotMatchErr},
		{TransferTxArg{Native: false, Decimals: 18,
			TxHash:       "0x20f7b501412aa33d47dd86b2b371b2db661a52363b42104eb60d87ffb3e89cc4",
			TokenAddress: "0x034029aFEf27c006D056Cc1eB98420aa0f3f8a1d",
			CheckArg: CheckArg{
				CheckFrom: true, From: "0xdf9B11c5Df465b288Af199E5291c254c81F2Ee56",
				CheckTo: true, To: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b88",
				CheckAmount: true, Amount: decimal.NewFromInt(10_000),
			},
		}, ToNotMatchErr},
		{TransferTxArg{Native: true, Decimals: 18,
			TxHash: "0xccbae6e82a12c38d1679cfb5e86ea0b8384960c330b53e91b18c0be87eeef596", TokenAddress: "",
			CheckArg: CheckArg{
				CheckFrom: true, From: "0xaa25Aa7a19f9c426E07dee59b12f944f4d9f1DD3",
				CheckTo: true, To: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89",
				CheckAmount: true, Amount: decimal.NewFromFloat(0.01),
			},
		}, AmountNotMatchErr},
		// 三种错误的忽略
		{TransferTxArg{Native: false, Decimals: 18,
			TxHash:       "0xdd09d28dbaa482d8af61bfa8c795d623713d2ae54099de4946b0cd8ed7684ee0",
			TokenAddress: "0x034029aFEf27c006D056Cc1eB98420aa0f3f8a1d",
			CheckArg: CheckArg{
				CheckFrom: false, From: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b88",
				CheckTo: true, To: "0xdf9B11c5Df465b288Af199E5291c254c81F2Ee56",
				CheckAmount: true, Amount: decimal.NewFromInt(100),
			},
		}, nil},
		{TransferTxArg{Native: false, Decimals: 18,
			TxHash:       "0x20f7b501412aa33d47dd86b2b371b2db661a52363b42104eb60d87ffb3e89cc4",
			TokenAddress: "0x034029aFEf27c006D056Cc1eB98420aa0f3f8a1d",
			CheckArg: CheckArg{
				CheckFrom: true, From: "0xdf9B11c5Df465b288Af199E5291c254c81F2Ee56",
				CheckTo: false, To: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b88",
				CheckAmount: true, Amount: decimal.NewFromInt(10_000),
			},
		}, nil},
		{TransferTxArg{Native: true, Decimals: 18,
			TxHash: "0xccbae6e82a12c38d1679cfb5e86ea0b8384960c330b53e91b18c0be87eeef596", TokenAddress: "",
			CheckArg: CheckArg{
				CheckFrom: true, From: "0xaa25Aa7a19f9c426E07dee59b12f944f4d9f1DD3",
				CheckTo: true, To: "0xf4062f23E362A8fd4FFb668Aacb78A24B1914b89",
				CheckAmount: false, Amount: decimal.NewFromFloat(0.01),
			},
		}, nil},
	}
	for _, testCase := range testCases {
		_, err := s.evmClient.GetTransferTx(s.ctx, testCase.TransferTxArg)
		assert.Equal(s.T(), testCase.err, errors.Cause(err))
	}
}

func (s *EVMClientTestSuite) TestTransfer() {
	usdtHash, err1 := s.evmClient.Transfer(s.ctx, TransferArg{
		PrivateKey:   s.privateKey,
		From:         s.from,
		Native:       false,
		ChainID:      s.chainID,
		Beneficiary:  "0x6c6988c875E998d678475D69D31a74000BAbBd57",
		TokenAddress: s.tokenAddress,
		Amount:       ToWei("1", 18),
	})
	assert.Nil(s.T(), err1)
	bnbHash, err2 := s.evmClient.Transfer(s.ctx, TransferArg{
		PrivateKey:   s.privateKey,
		From:         s.from,
		Native:       true,
		ChainID:      s.chainID,
		Beneficiary:  "0x6c6988c875E998d678475D69D31a74000BAbBd57",
		TokenAddress: "",
		Amount:       ToWei("0.01", 18),
	})
	assert.Nil(s.T(), err2)
	s.T().Log("USDT hash: ", usdtHash)
	s.T().Log("BNB hash: ", bnbHash)
}

func (s *EVMClientTestSuite) TestSplitTransfer() {
	// ERC20 转账
	txBytes, err1 := s.evmClient.GenerateTransferData(s.ctx, TransferArg{
		From:         s.from,
		Native:       false,
		ChainID:      s.chainID,
		Beneficiary:  "0x6c6988c875E998d678475D69D31a74000BAbBd57",
		TokenAddress: s.tokenAddress,
		Amount:       ToWei("1", 18),
	})
	assert.Nil(s.T(), err1)
	signBytes, err2 := s.evmClient.SignTransfer(s.ctx, SignTransferArg{PrivateKey: s.privateKey, ChainID: s.chainID, TxBytes: txBytes})
	assert.Nil(s.T(), err2)
	hash, err := s.evmClient.SendTransactionManually(s.ctx, signBytes)
	assert.Nil(s.T(), err)
	s.T().Log("USDT hash: ", hash)

	// 原声币种转账
	txBytes, err1 = s.evmClient.GenerateTransferData(s.ctx, TransferArg{
		From:         s.from,
		Native:       true,
		ChainID:      s.chainID,
		Beneficiary:  "0x6c6988c875E998d678475D69D31a74000BAbBd57",
		TokenAddress: "",
		Amount:       ToWei("0.01", 18),
	})
	assert.Nil(s.T(), err1)
	signBytes, err2 = s.evmClient.SignTransfer(s.ctx, SignTransferArg{PrivateKey: s.privateKey, ChainID: s.chainID, TxBytes: txBytes})
	assert.Nil(s.T(), err2)
	hash, err = s.evmClient.SendTransactionManually(s.ctx, signBytes)
	assert.Nil(s.T(), err)
	s.T().Log("USDT hash: ", hash)
}

func (s *EVMClientTestSuite) TestApprove() {
	err := s.evmClient.Approve(s.ctx, ApproveArg{
		PrivateKey:   s.privateKey,
		ChainID:      s.chainID,
		TokenAddress: s.tokenAddress,
		Spender:      "0x6c6988c875E998d678475D69D31a74000BAbBd57",
		Amount:       ToWei("100", defaultEvmDecimals),
	})
	assert.Nil(s.T(), err)
}

func (s *EVMClientTestSuite) TestAllowance() {
	allowance, err := s.evmClient.Allowance(s.ctx, AllowanceArg{
		TokenAddress: s.tokenAddress,
		Decimals:     defaultEvmDecimals,
		Owner:        s.from,
		Spender:      "0x6c6988c875E998d678475D69D31a74000BAbBd57",
	})
	assert.Nil(s.T(), err)
	s.T().Log(allowance.String())
}

func (s *EVMClientTestSuite) TestTransferFrom() {
	// 0x6c6988c875E998d678475D69D31a74000BAbBd57 的私钥
	spenderPrivateKey := "c83c6f680903b488702b27539d490a4a2a47ea1c8d1c4b1b5e7c4a2ff87448f9"
	hash, err := s.evmClient.TransferFrom(s.ctx, TransformArg{
		PrivateKey:   spenderPrivateKey,
		From:         s.from,
		Beneficiary:  "0xdf9B11c5Df465b288Af199E5291c254c81F2Ee56",
		TokenAddress: s.tokenAddress,
		ChainID:      s.chainID,
		Amount:       ToWei("0.1", defaultEvmDecimals),
	})
	assert.Nil(s.T(), err)
	s.T().Log(hash)
}
