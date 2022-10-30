package command

import "fmt"

// Doctor ...
type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// ICommand 抽象命令
type ICommand interface {
	Execute()
}

// TreatEyeCmd ...
type TreatEyeCmd struct {
	doctor *Doctor // 可以组合进来，也可以将其作为Execute函数的入参数传入，可根据实际情况采用不同方式
}

// Execute ...
func (te *TreatEyeCmd) Execute() {
	te.doctor.treatEye()
}

// TreatNoseCmd ..
type TreatNoseCmd struct {
	doctor *Doctor
}

// Execute ..
func (tn *TreatNoseCmd) Execute() {
	tn.doctor.treatNose()
}

// Nurse ...
type Nurse struct {
	CmdList []ICommand
}

// Notify ...
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.Execute()
	}
}

func logic() {
	//依赖病单，通过填写病单，让医生看病
	doctor := new(Doctor)
	//治疗眼睛的病单
	cmdEye := TreatEyeCmd{doctor}
	//治疗鼻子的病单
	cmdNose := TreatNoseCmd{doctor}

	//护士
	nurse := new(Nurse)
	//收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)

	//执行病单指令
	nurse.Notify()
}

// Index ...
func Index() {
	logic()
}
