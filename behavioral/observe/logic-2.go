package observe

import (
	"fmt"
	"time"
)

const (
	// PGaiBang ..
	PGaiBang string = "丐帮"
	// PMingJiao ...
	PMingJiao string = "明教"
)

// IListener2 ..
type IListener2 interface {
	// 当同伴被揍了该怎么办
	Fight(another IListener2, baixiao INotifier2)
	OnFriendBeFight(event *Event)
	GetName() string
	GetParty() string
	Title() string
}

// INotifier2 ...
type INotifier2 interface {
	AddListener(listener IListener2)
	RemoveListener(listener IListener2)
	Notify(event *Event)
}

// Event ...
type Event struct {
	Noti    INotifier2 // 被知晓的通知者
	One     IListener2 // 事件主动发出者
	Another IListener2 // 时间被动接收者
	Msg     string     // 具体消息
}

//====实现层====

// Hero ...
type Hero struct {
	Name  string
	Party string
}

// GetName ...
func (h *Hero) GetName() string {
	return h.Name
}

// GetParty ...
func (h *Hero) GetParty() string {
	return h.Party
}

// Title ...
func (h *Hero) Title() string {
	return fmt.Sprintf("[%s]:%s", h.Party, h.Name)
}

// Fight ...
func (h *Hero) Fight(another IListener2, baixiao INotifier2) {
	msg := fmt.Sprintf("%s 将 %s 揍了...", h.Title(), another.Title())

	event := new(Event)
	event.Noti = baixiao
	event.One = h
	event.Another = another
	event.Msg = msg

	baixiao.Notify(event)
}

// OnFriendBeFight ...
func (h *Hero) OnFriendBeFight(event *Event) {
	// 当事人双方不做出任何反应
	if h.Name == event.One.GetName() || h.Name == event.Another.GetName() {
		return
	}

	// 本帮派同伴将其他门派揍了，要拍手叫好!
	if h.GetParty() == event.One.GetParty() {
		fmt.Println(h.Title(), "得知消息，拍手叫好！！！")
		return
	}

	// 本帮派同伴被其他门派揍了，要主动报仇反击!
	if h.GetParty() == event.Another.GetParty() {
		fmt.Println(h.Title(), "得知消息，发起报仇反击！！！")
		time.Sleep(5 * time.Second)
		h.Fight(event.One, event.Noti)
	}
}

// BaiXiao ...
type BaiXiao struct {
	heroList []IListener2
}

// AddListener ...
func (b *BaiXiao) AddListener(listener IListener2) {
	b.heroList = append(b.heroList, listener)
}

// RemoveListener ...
func (b *BaiXiao) RemoveListener(listener IListener2) {
	for index, l := range b.heroList {
		//找到要删除的元素位置
		if listener == l {
			//将删除的点前后的元素链接起来
			b.heroList = append(b.heroList[:index], b.heroList[index+1:]...)
			break
		}
	}
}

// Notify ...
func (b *BaiXiao) Notify(event *Event) {
	fmt.Println("======================")
	fmt.Println("【世界消息】 百晓生广播消息: ", event.Msg)
	for _, l := range b.heroList {
		go l.OnFriendBeFight(event)
	}
}

func logic2() {
	var h1 IListener2
	h1 = &Hero{
		"黄蓉",
		PGaiBang,
	}

	var h2 IListener2
	h2 = &Hero{
		"洪七公",
		PGaiBang,
	}

	var h3 IListener2
	h3 = &Hero{
		"乔峰",
		PGaiBang,
	}

	var h4 IListener2
	h4 = &Hero{
		"张无忌",
		PMingJiao,
	}

	var h5 IListener2
	h5 = &Hero{
		"韦一笑",
		PMingJiao,
	}

	var h6 IListener2
	h6 = &Hero{
		"金毛狮王",
		PMingJiao,
	}

	var baixiao INotifier2
	baixiao = new(BaiXiao)

	baixiao.AddListener(h1)
	baixiao.AddListener(h2)
	baixiao.AddListener(h3)
	baixiao.AddListener(h4)
	baixiao.AddListener(h5)
	baixiao.AddListener(h6)

	fmt.Println("武林一片平静.....直到")
	h1.Fight(h5, baixiao)

	select {}
}
