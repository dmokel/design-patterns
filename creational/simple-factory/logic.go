package simplefactory

import "fmt"

// Fruit ...
type Fruit struct {
	// some fields
	Name string
}

// Show ...
func (f *Fruit) Show() {
	if f.Name == "apple" {
		fmt.Println("this is apple")
	} else if f.Name == "pear" {
		fmt.Println("this is pear")
	}
}

// NewFruit ...
func NewFruit(name string) *Fruit {
	fruit := new(Fruit)

	if name == "apple" {
		// set the apple fields
		fruit.Name = "apple"
	} else if name == "pear" {
		// set the pear fields
		fruit.Name = "pear"
	}

	return fruit
}

func logicNon() {
	// get an apple
	apple := NewFruit("apple")
	apple.Show()

	// get a pear
	pear := NewFruit("pear")
	pear.Show()
}

// simple factory
//
//
//

// IFruit ...
type IFruit interface {
	Show()
}

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

// Factory ...
type Factory struct{}

// CreateFruit ...
func (f *Factory) CreateFruit(name string) IFruit {
	var fruit IFruit

	if name == "apple" {
		fruit = &Apple{
			Name: "apple",
		}
	} else if name == "pear" {
		fruit = &Pear{
			Name: "pear",
		}
	}

	return fruit
}

func logicWithSampleFactory() {
	// first new an factory
	factory := new(Factory)

	// get an apple
	apple := factory.CreateFruit("apple")
	apple.Show()

	// get an pear
	pear := factory.CreateFruit("pear")
	pear.Show()
}

// Index ...
func Index() {
	logicNon()
	fmt.Println("======")
	logicWithSampleFactory()
}
