package singleton

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*三个要点：
 * 一是某个类只能有一个实例；
 * 二是它必须自行创建这个实例；
 * 三是它必须自行向整个系统提供这个实例。
 */

// 该类只能有一个实例
// 意味着该类不能被外部直接获取
type singleton struct{}

// 自行创建该实例
// 意味着必须有一个实例，即有一个指针变量指向该实例，且该指针不能被外部所修改，
var instance *singleton = new(singleton)

// 向系统提供该实例，且系统不能直接修改具体实例，即不能直接修改该指针变量所存储的地址
// 因此，向系统提供一份拷贝的指针变量，外部修改的只是一个外部指针变量，无法修改包内的指针变量

// GetInstace ...
func GetInstace() *singleton {
	return instance
}

func (s *singleton) Show() {
	fmt.Println("i am singleton")
}

func logic() {
	var instance = GetInstace()
	instance.Show()

	instance = nil

	var instance2 = GetInstace()
	instance2.Show() // 不影响instance2
}

//====lazy init====

var instance2 *singleton
var once sync.Once

// GetInstace2 ...
func GetInstace2() *singleton {
	once.Do(func() {
		instance2 = new(singleton)
	})
	return instance2
}

var lock sync.Locker
var initialized uint32

// GetInstace3 ...
func GetInstace3() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance2
	}

	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		instance2 = new(singleton)
		atomic.StoreUint32(&initialized, 1)
	}

	return instance2
}

// Index ...
func Index() {
	logic()
}
