package facade

import "fmt"

type subSystemA struct{}

func (a *subSystemA) MethodA() {
	fmt.Println("subSystemA MethodA")
}

type subSystemB struct{}

func (b *subSystemB) MethodB() {
	fmt.Println("subSystemB MethodB")
}

type subSystemC struct{}

func (c *subSystemC) MethodC() {
	fmt.Println("subSystemC MethodC")
}

type subSystemD struct{}

func (d *subSystemD) MethodD() {
	fmt.Println("subSystemD MethodD")
}

// Facade ...
type Facade struct {
	a *subSystemA
	b *subSystemB
	c *subSystemC
	d *subSystemD
}

// MethodOne ...
func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

// MethodTwo ...
func (f *Facade) MethodTwo() {
	f.c.MethodC()
	f.d.MethodD()
}

func logic() {
	// 不用外观模式调用MethodA和MethodB
	sa := new(subSystemA)
	sb := new(subSystemB)
	sa.MethodA()
	sb.MethodB()

	// 用外观模式调用MethodA何MethodB
	f := Facade{
		a: new(subSystemA),
		b: new(subSystemB),
		c: new(subSystemC),
		d: new(subSystemD),
	}
	f.MethodOne()
}

// 家庭设备管理案例

// TV 电视机
type TV struct{}

// On ...
func (t *TV) On() {
	fmt.Println("打开 电视机")
}

// Off ...
func (t *TV) Off() {
	fmt.Println("关闭 电视机")
}

// VoiceBox 电视机
type VoiceBox struct{}

// On ...
func (v *VoiceBox) On() {
	fmt.Println("打开 音箱")
}

// Off ...
func (v *VoiceBox) Off() {
	fmt.Println("关闭 音箱")
}

// Light 灯光
type Light struct{}

// On ...
func (l *Light) On() {
	fmt.Println("打开 灯光")
}

// Off ...
func (l *Light) Off() {
	fmt.Println("关闭 灯光")
}

// Xbox 游戏机
type Xbox struct{}

// On ...
func (x *Xbox) On() {
	fmt.Println("打开 游戏机")
}

// Off ...
func (x *Xbox) Off() {
	fmt.Println("关闭 游戏机")
}

// MicroPhone 麦克风
type MicroPhone struct{}

// On ...
func (m *MicroPhone) On() {
	fmt.Println("打开 麦克风")
}

// Off ...
func (m *MicroPhone) Off() {
	fmt.Println("关闭 麦克风")
}

// Projector 投影仪
type Projector struct{}

// On ...
func (p *Projector) On() {
	fmt.Println("打开 投影仪")
}

// Off ...
func (p *Projector) Off() {
	fmt.Println("关闭 投影仪")
}

// HomePlayerFacade 家庭影院(外观)
type HomePlayerFacade struct {
	tv    TV
	vb    VoiceBox
	light Light
	xbox  Xbox
	mp    MicroPhone
	pro   Projector
}

// DoKTV KTV模式
func (hp *HomePlayerFacade) DoKTV() {
	fmt.Println("家庭影院进入KTV模式")
	hp.tv.On()
	hp.pro.On()
	hp.mp.On()
	hp.light.Off()
	hp.vb.On()
}

// DoGame 游戏模式
func (hp *HomePlayerFacade) DoGame() {
	fmt.Println("家庭影院进入Game模式")
	hp.tv.On()
	hp.light.On()
	hp.xbox.On()
}

func logic2() {
	homePlayer := new(HomePlayerFacade) // HomePlayerFacade组合进来的子组件不是指针类型，所以默认不是nil，而是自动初始化的结构体

	homePlayer.DoKTV()

	fmt.Println("------------")

	homePlayer.DoGame()
}

// Index ...
func Index() {
	logic()

	fmt.Println("======")

	logic2()
}
