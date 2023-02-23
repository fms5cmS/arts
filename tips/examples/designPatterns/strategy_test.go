package designPatterns

import (
	"fmt"
	"testing"
)

// 上下文维护指向具体策略的引用，上下文不执行任务，而是将工作委派给已连接的策略对象
type PayCtx struct {
	// 提供支付能力的接口实现
	payBehavior PayBehavior
	// 支付参数
	payParams map[string]interface{}
}

func (px *PayCtx) setPayBehavior(p PayBehavior) {
	px.payBehavior = p
}

func (px *PayCtx) Pay() {
	px.payBehavior.OrderPay(px)
}

func NewPayCtx(p PayBehavior) *PayCtx {
	// 支付参数，Mock数据
	params := map[string]interface{}{
		"appId": "234fdfdngj4",
		"mchId": 123456,
	}
	return &PayCtx{
		payBehavior: p,
		payParams:   params,
	}
}

// 抽象策略
type PayBehavior interface {
	OrderPay(px *PayCtx)
}

// 具体的策略实现一：Wx 支付
type WxPay struct{}

func (*WxPay) OrderPay(px *PayCtx) {
	fmt.Printf("Wx支付加工支付请求 %v\n", px.payParams)
	fmt.Println("正在使用Wx支付进行支付")
}

// 具体的策略实现二：三方支付
type ThirdPay struct{}

func (*ThirdPay) OrderPay(px *PayCtx) {
	fmt.Printf("三方支付加工支付请求 %v\n", px.payParams)
	fmt.Println("正在使用三方支付进行支付")
}

func TestStrategy(t *testing.T) {
	wxPay := &WxPay{}
	px := NewPayCtx(wxPay)
	px.Pay()
	// 假设现在发现微信支付没钱，改用三方支付进行支付
	thPay := &ThirdPay{}
	px.setPayBehavior(thPay)
	px.Pay()
}
