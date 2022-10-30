package templatemethod

import "fmt"

// 模板方法
// 业务方只关心宏观结果，不关心底层模块具体方法的实现和方法的组合调用

// IBeverage ...
type IBeverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()

	// hooks钩子，控制模板类内部逻辑流
	WantAddThings() bool
}

// 模板类，组合抽象类，并基于抽象类定义的方法，预写算法
type template struct {
	b IBeverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

// MakeCoffee ..
type MakeCoffee struct {
	template // 继承template，主要是继承template的属性和模板方法
}

// NewMakeCoffee ...
func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)

	makeCoffee.b = makeCoffee

	return makeCoffee
}

// BoilWater ...
func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

// Brew ...
func (mc *MakeCoffee) Brew() {
	fmt.Println("用水冲咖啡豆")
}

// PourInCup ...
func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将充好的咖啡倒入陶瓷杯中")
}

// AddThings ...
func (mc *MakeCoffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

// WantAddThings ...
func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

// MakeTea ...
type MakeTea struct {
	template
}

// NewMakeTea ...
func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)

	makeTea.b = makeTea

	return makeTea
}

// BoilWater ...
func (mt *MakeTea) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

// Brew ...
func (mt *MakeTea) Brew() {
	fmt.Println("用水冲茶叶")
}

// PourInCup ...
func (mt *MakeTea) PourInCup() {
	fmt.Println("将充好的咖啡倒入茶壶中")
}

// AddThings ...
func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

// WantAddThings ...
func (mt *MakeTea) WantAddThings() bool {
	return false
}

func logic() {
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBeverage()

	fmt.Println("========")

	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}

// Index ...
func Index() {
	logic()
}

// 优缺点
// 优点：
// 1、父类/模板类通过组合抽象类，抽象类定义了子步骤，并在父类/模板类中形式化的定义一个算法，而由其子类实现子步骤的具体逻辑，
// 在子类详细实现算法具体步骤的逻辑时，其不会改变算法步骤的执行次序。
