package adapter

import "fmt"

// IV5 适配的目标
type IV5 interface {
	UseV5()
}

// Phone 业务类，依赖5V接口
type Phone struct {
	v5 IV5
}

// Charge ...
func (p *Phone) Charge() {
	fmt.Println("phone use 5v charge")
	p.v5.UseV5()
}

// NewPhone ...
func NewPhone(v5 IV5) *Phone {
	return &Phone{v5}
}

// V220 现有组件
type V220 struct{}

// UseV220 ...
func (v220 V220) UseV220() {
	fmt.Println("use 220v")
}

// Adapter ...
type Adapter struct {
	v220 *V220
}

// UseV5 ...
func (a *Adapter) UseV5() {
	fmt.Println("adapter use 220v")
	a.v220.UseV220()
}

// NewAdapter ...
func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func phonelogic() {
	// 在不修改phone组件前提下（oop开闭原则），又必须使用现有220v组件（该现有组件也是不可修改的），则需要一个适配器
	phone := NewPhone(NewAdapter(new(V220)))
	phone.Charge()
}
