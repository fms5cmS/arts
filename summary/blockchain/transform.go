package blockchain

import (
	"github.com/shopspring/decimal"
	"math/big"
)

// ToWei 将金额转为最小单位
func ToWei(iAmount interface{}, decimals int32) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iAmount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	case *big.Int:
		amount = decimal.NewFromBigInt(v, 0)
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)
	wei := new(big.Int)
	wei.SetString(result.String(), 10)
	return wei
}

// ToDecimal 将最小单位的金额转为常用单位
func ToDecimal(iValue interface{}, decimals int32) decimal.Decimal {
	value := new(big.Int)
	switch v := iValue.(type) {
	case string:
		value.SetString(v, 10)
	case decimal.Decimal:
		value = v.BigInt()
	case int:
		value = big.NewInt(int64(v))
	case *big.Int:
		value = v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)
	return result
}
