package abstractfactory

import "fmt"

// 在工厂方法的案例中，已有两个具体的工厂类，分别为appleFactory和pearFactory，
// 如果我们有以下扩展需求
// 1、中国的苹果，中国的梨，
// 2、日本的苹果，日本的梨，
// 3、美国的苹果，美国的梨，
// 如果按照工厂方法的设计模式，每一种都新增一个新的工厂类，会导致工厂类数量急剧增多
// 而正如我们文字描述一样，该拓展需求具有一定的规律性，为苹果/梨与中国/日本/美国的排列组合
// 一者称为产品等级，一者称为产品族
// 产品等级具有一定的稳定性，即长时间内某领域内的产品等级划分是不变的，
// 产品族具有可变性、灵活性，即可以根据需求变化随时调整产品族
// 因此，上述情况中，
// 苹果/梨称为产品等级，我们假设水果产品领域只有苹果和梨两种，第三种需要经过较长期的科学研究与栽培才能培育出来，所以具有稳定性。
// 中国/日本/美国称为产品族，只要第四个国家引入苹果和梨的种植技术，那么将会出现该国家的苹果和梨，所以具有可变形、灵活性。
// 就如同设计好的主板具有几个模块是固定不变的一般，而生产该主板的厂商可以处于快速的变化中。
//
// 综上，提出抽象工厂模式
// 将不同产品族具化为不同的产品族工厂，该工厂对外提供一个产品族内所有产品等级的实例的创建

//====抽象层====

// IApple ...
type IApple interface {
	Show()
}

// IPear ...
type IPear interface {
	Show()
}

// IFactory ...
type IFactory interface {
	CreateApple() IApple
	CreatePear() IPear
}

//====实现层====

// 中国的产品族

// ChinaApple ...
type ChinaApple struct{}

// Show ...
func (ca *ChinaApple) Show() {
	fmt.Println("i am china apple")
}

// ChinaPear ...
type ChinaPear struct{}

// Show ...
func (cp *ChinaPear) Show() {
	fmt.Println("i am china pear")
}

// ChinaFactory ...
type ChinaFactory struct{}

// CreateApple ...
func (cf *ChinaFactory) CreateApple() IApple {
	var apple IApple
	apple = new(ChinaApple)
	return apple
}

// CreatePear ...
func (cf *ChinaFactory) CreatePear() IPear {
	var pear IPear
	pear = new(ChinaPear)
	return pear
}

//美国的产品族

// AmericaApple ...
type AmericaApple struct{}

// Show ...
func (aa *AmericaApple) Show() {
	fmt.Println("i am america apple")
}

// AmericaPear ...
type AmericaPear struct{}

// Show ...
func (ap *AmericaPear) Show() {
	fmt.Println("i am america pear")
}

// AmericaFactory ...
type AmericaFactory struct{}

// CreateApple ...
func (af *AmericaFactory) CreateApple() IApple {
	var apple IApple
	apple = new(AmericaApple)
	return apple
}

// CreatePear ...
func (af *AmericaFactory) CreatePear() IPear {
	var pear IPear
	pear = new(AmericaPear)
	return pear
}

// Index ...
func Index() {
	var factory IFactory

	// 中国苹果，中国梨
	factory = new(ChinaFactory)
	cApple := factory.CreateApple()
	cPear := factory.CreatePear()
	cApple.Show()
	cPear.Show()

	// 美国苹果，美国梨
	factory = new(AmericaFactory)
	aApple := factory.CreateApple()
	aPear := factory.CreatePear()
	aApple.Show()
	aPear.Show()
}
