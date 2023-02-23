package designPatterns

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 接口或抽象类声明方法
type BankBusinessHandler interface {
	// 取号
	TakeRowNumber()
	// 排队
	WaitInHead()
	// 处理业务
	HandleBusiness()
	// 对服务做出评价
	Commentate()
	// 用于在流程中判断是否为 VIP
	// 钩子（可选步骤）
	CheckVipIdentity() bool
}

// Go 不支持抽象类和继承，通过嵌套来实现
type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

// 模版方法
func (b *BankBusinessExecutor) ExecuteBusiness() {
	b.handler.TakeRowNumber()
	if !b.handler.CheckVipIdentity() {
		b.handler.WaitInHead()
	}
	b.handler.HandleBusiness()
	b.handler.Commentate()
}

func NewBankBusinessExecutor(businessHandler BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{handler: businessHandler}
}

/**           通用实现                 */
type DefaultBusinessHandler struct{}

func (*DefaultBusinessHandler) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) + " ，注意排队情况，过号后顺延三个安排")
}

func (dbh *DefaultBusinessHandler) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(5 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (*DefaultBusinessHandler) Commentate() {

	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

func (*DefaultBusinessHandler) CheckVipIdentity() bool {
	// 留给具体实现类实现
	return false
}

/**                         特有实现               */

type DepositBusinessHandler struct {
	*DefaultBusinessHandler // 嵌入一个默认的实现，自己仅需实现特有的逻辑即可
	userVip                 bool
}

func (*DepositBusinessHandler) HandleBusiness() {
	fmt.Println("账户存储很多万人民币...")
}

func (dh *DepositBusinessHandler) CheckVipIdentity() bool {
	return dh.userVip
}

func TestTemplate(t *testing.T) {
	dh := &DepositBusinessHandler{userVip: false}
	e := NewBankBusinessExecutor(dh)
	e.ExecuteBusiness()
}
