package theory

import "fmt"

// BMW ..
type BMW struct{}

// Running ...
func (b *BMW) Running() {
	fmt.Println("BMW is Running")
}

// BYD ..
type BYD struct{}

// Running ...
func (b *BYD) Running() {
	fmt.Println("BYD is Running")
}

// ZS ...
type ZS struct{}

// DriveBMW ...
func (zs *ZS) DriveBMW(bmw *BMW) {
	fmt.Println("zs drive bmw")
	bmw.Running()
}

// DriveBYD ...
func (zs *ZS) DriveBYD(byd *BYD) {
	fmt.Println("zs drive BYD")
	byd.Running()
}

// WW ...
type WW struct{}

// DriveBMW ...
func (ww *WW) DriveBMW(bmw *BMW) {
	fmt.Println("ww drive bmw")
	bmw.Running()
}

// DriveBYD ...
func (ww *WW) DriveBYD(byd *BYD) {
	fmt.Println("ww drive byd")
	byd.Running()
}

func logicNonDip() {
	bmw := new(BMW)
	byd := new(BYD)

	zs := new(ZS)
	zs.DriveBMW(bmw)
	zs.DriveBYD(byd)

	ww := new(WW)
	ww.DriveBMW(bmw)
	ww.DriveBYD(byd)
}

// 具体的司机和具体的汽车之间并不是有强关联的实体，
// ta们可以相互独立的拓宽自己的实体种类，因此ta们不属于适合高内聚的两个模块，而更适合为两个低耦合的模块
// 如何解耦：通过抽象出两个模块的接口层/抽象层，且接口层的定义只依赖对方的接口层，从而两个模块的实现层也只依赖对方的接口层；
// 同时业务层实体定义也只依赖两个模块的接口层

// ====抽象层====

// Car ...
type Car interface {
	Running()
	GetName() string
}

// Driver ...
type Driver interface {
	Drive(car Car)
}

// ====实现层====

// BMW2 ...
type BMW2 struct {
	Name string
}

// Running ...
func (bmw2 BMW2) Running() {
	fmt.Println("BMW2 is Running")
}

// GetName ...
func (bmw2 BMW2) GetName() string {
	return bmw2.Name
}

// BYD2 ...
type BYD2 struct {
	Name string
}

// Running ...
func (byd2 BYD2) Running() {
	fmt.Println("BYD2 is Running")
}

// GetName ...
func (byd2 BYD2) GetName() string {
	return byd2.Name
}

// ZS2 ...
type ZS2 struct{}

// Drive ...
func (zs2 ZS2) Drive(car Car) {
	fmt.Println("ZS2 drive ", car.GetName())
	car.Running()
}

// WW2 ...
type WW2 struct{}

// Drive ...
func (ww2 WW2) Drive(car Car) {
	fmt.Println("WW2 drive ", car.GetName())
	car.Running()
}

func logicDip() {
	var bmw2 Car
	bmw2 = &BMW2{
		Name: "BMW2",
	}

	var byd2 Car
	byd2 = &BYD2{
		Name: "BYD2",
	}

	var ww2 Driver
	ww2 = new(WW2)
	ww2.Drive(bmw2)
	ww2.Drive(byd2)

	var zs2 Driver
	zs2 = new(ZS2)
	zs2.Drive(bmw2)
	zs2.Drive(byd2)
}
