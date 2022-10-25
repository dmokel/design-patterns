package theory

import "fmt"

// Index ...
func Index() {
	// ocp
	logicNonOcp()
	fmt.Println("======")
	logicOcp()
	fmt.Println("======")
	logic()
	fmt.Println("======")

	// dip
	logicNonDip()
	fmt.Println("======")
	logicDip()
	fmt.Println("======")
}
