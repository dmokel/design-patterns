package observe

import "fmt"

//====抽象层====

// IListener ...
type IListener interface {
	OnTeacherComming()
}

// INotifier ...
type INotifier interface {
	AddListener(listener IListener)
	RemoveListener(listener IListener)
	Notify()
}

//====实现层====

// StuZS ...
type StuZS struct {
	Badthing string
}

// OnTeacherComming ...
func (zs *StuZS) OnTeacherComming() {
	fmt.Println("张三 停止", zs.Badthing)
}

// DoBadThing ...
func (zs *StuZS) DoBadThing() {
	fmt.Println("张三 开始", zs.Badthing)
}

// StuWW ...
type StuWW struct {
	Badthing string
}

// OnTeacherComming ...
func (ww *StuWW) OnTeacherComming() {
	fmt.Println("王五 停止", ww.Badthing)
}

// DoBadThing ...
func (ww *StuWW) DoBadThing() {
	fmt.Println("王五 开始", ww.Badthing)
}

// Monitor ....
type Monitor struct {
	ListenerList []IListener
}

// AddListener ...
func (m *Monitor) AddListener(listener IListener) {
	m.ListenerList = append(m.ListenerList, listener)
}

// RemoveListener ...
func (m *Monitor) RemoveListener(listener IListener) {
	for index, l := range m.ListenerList {
		if l == listener {
			m.ListenerList = append(m.ListenerList[:index], m.ListenerList[index+1:]...)
			break
		}
	}
}

// Notify ...
func (m *Monitor) Notify() {
	for _, l := range m.ListenerList {
		l.OnTeacherComming()
	}
}

func logic() {
	zs := &StuZS{
		Badthing: "抄作业",
	}
	ww := &StuWW{
		Badthing: "打游戏",
	}

	monitor := new(Monitor)
	monitor.AddListener(zs)
	monitor.AddListener(ww)

	fmt.Println("上课了，但是老师没有来，学生们都在忙自己的事...")
	zs.DoBadThing()
	ww.DoBadThing()

	fmt.Println("这时候老师来了，班长给学什么使了一个眼神...")
	monitor.Notify()
}

// Index ...
func Index() {
	// logic()

	logic2()
}
