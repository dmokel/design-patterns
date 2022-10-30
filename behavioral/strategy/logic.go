package strategy

import "fmt"

// IWeaponStrategy 武器策略(抽象的策略)
type IWeaponStrategy interface {
	UseWeapon() //使用武器
}

// Ak47 具体的策略
type Ak47 struct{}

// UseWeapon ...
func (ak *Ak47) UseWeapon() {
	fmt.Println("使用Ak47 去战斗")
}

// Knife 具体的策略
type Knife struct{}

// UseWeapon ...
func (k *Knife) UseWeapon() {
	fmt.Println("使用匕首 去战斗")
}

// Hero 环境类
type Hero struct {
	strategy IWeaponStrategy //拥有一个抽象的策略
}

// SetWeaponStrategy 设置一个策略
func (h *Hero) SetWeaponStrategy(s IWeaponStrategy) {
	h.strategy = s
}

// Fight ...
func (h *Hero) Fight() {
	h.strategy.UseWeapon() //调用策略
}

func logic() {
	hero := new(Hero)

	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}

// Index ...
func Index() {
	logic()
}
