package factorymethod

import "fmt"

//====抽象层====

// IFruit ...
type IFruit interface {
	Show()
}

// IFactory ...
type IFactory interface {
	CreateFruit() IFruit
}

//====实现层====

// Apple ...
type Apple struct {
	Name string
}

// Show ...
func (a *Apple) Show() {
	fmt.Println("this is ", a.Name)
}

// Pear ...
type Pear struct {
	Name string
}

// Show ...
func (p *Pear) Show() {
	fmt.Println("this is ", p.Name)
}

// AppleFactory ...
type AppleFactory struct{}

// CreateFruit ...
func (af *AppleFactory) CreateFruit() IFruit {
	var apple IFruit

	apple = &Apple{
		Name: "apple",
	}

	return apple
}

// PearFactory ...
type PearFactory struct{}

// CreateFruit ...
func (pf *PearFactory) CreateFruit() IFruit {
	var pear IFruit

	pear = &Pear{
		Name: "pear",
	}

	return pear
}

// Index ...
func Index() {
	var factory IFactory

	// apple
	factory = new(AppleFactory)
	apple := factory.CreateFruit()
	apple.Show()

	// pear
	factory = new(PearFactory)
	pear := factory.CreateFruit()
	pear.Show()
}
