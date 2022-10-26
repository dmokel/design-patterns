package proxy

import "fmt"

// Goods ...
type Goods struct {
	Kind string
	Fact bool
}

//====抽象层====

// IShopping ...
type IShopping interface {
	Buy(goods *Goods)
}

//====实现层====

// KoreaShopping ...
type KoreaShopping struct{}

// Buy ...
func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物，买了 ", goods.Kind)
}

// AmericanShopping ...
type AmericanShopping struct{}

// Buy ...
func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物，买了 ", goods.Kind)
}

// 满足开闭原则，通过一个代理购物类来为原购物类新增相关功能，
// 并控制业务方只与代理类进行交互，由代理类来完成与原购物类的交互，而业务方不直接与原购物类交互
// 即代理类形成了对原购物类的访问控制

// ShoppingProxy ...
type ShoppingProxy struct {
	shopping IShopping
}

// NewShoppingProxy ...
func NewShoppingProxy(shopping IShopping) IShopping {
	return &ShoppingProxy{
		shopping: shopping,
	}
}

// Buy ...
// 需要实现原抽象购物类的接口，在该方法中完成前处理、原处理、后处理逻辑，
// 且对业务方而言影响最小，依旧保持原有的方法调用逻辑，同时也依旧只依赖于原抽象购物类
func (sp *ShoppingProxy) Buy(goods *Goods) {
	if sp.distinguish(goods) == true {
		sp.shopping.Buy(goods)
		sp.check(goods)
	}
}

// 验货
func (sp *ShoppingProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买。")
	}
	return goods.Fact
}

// 安检
func (sp *ShoppingProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}

func logic() {
	var shopping IShopping

	shopping = NewShoppingProxy(new(KoreaShopping))
	shopping.Buy(&Goods{
		Kind: "化妆品",
		Fact: true,
	})
	shopping.Buy(&Goods{
		Kind: "cet4证书",
		Fact: false,
	})
}

// Index ...
func Index() {
	logic()
}
