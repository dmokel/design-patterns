package decorator

import "fmt"

// 代理模式中，一般没有代理基础类（抽象层），由具体的代理类组合抽象构件，并实现抽象层定义的方法，增加一些业务逻辑
// 所以，代理模式从概念和实现上更倾向于封装并控制业务方对一个基础构件的访问使用，
// 而不像装饰器模式这般具有更灵活的组合方式，从而形成新构件

//====抽象层====

// IPhone ...
type IPhone interface {
	Show()
}

// Decorator ...
// 装饰器基础类，该类本应该为interface，但是golang interface语法不可以有成员属性
type Decorator struct {
	phone IPhone
}

// Show ...这个show方法是否一定要实现?
// 经验证，该show方法不需要实现也是可以正常使用的，因为具体的装饰器类主要继承该基础类的phone属性
func (d *Decorator) Show() {}

//====实现层====

// Huawei ...
type Huawei struct{}

// Show ...
func (hw *Huawei) Show() {
	fmt.Println("华为手机")
}

// Xiaomi ...
type Xiaomi struct{}

// Show ...
func (xm *Xiaomi) Show() {
	fmt.Println("小米手机")
}

// 具体的装饰器类

// MoDecorator 膜装饰器
type MoDecorator struct {
	Decorator // 继承基础装饰器类，主要是继承phone属性
}

// Show ...
func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("贴膜的手机")
}

// NewMoDecorator ...
func NewMoDecorator(phone IPhone) IPhone {
	return &MoDecorator{Decorator{phone}}
}

// KeDecorator 壳装饰器
type KeDecorator struct {
	Decorator // 继承基础类的phone属性
}

// Show ...
func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("套壳的手机")
}

// NewKeDecorator ...
func NewKeDecorator(phone IPhone) IPhone {
	return &KeDecorator{Decorator{phone}}
}

func logic() {
	var huawei IPhone

	huawei = new(Huawei)
	huawei.Show()

	fmt.Println("=====")
	// 利用膜装饰器得到新构件
	moHuawei := NewMoDecorator(huawei)
	moHuawei.Show()

	fmt.Println("=====")
	// 利用壳装饰器得到新构件
	keHuawei := NewKeDecorator(huawei)
	keHuawei.Show()

	fmt.Println("=====")
	kemoHuawei := NewKeDecorator(NewMoDecorator(huawei))
	kemoHuawei.Show()
}

// Index ...
func Index() {
	logic()
}
