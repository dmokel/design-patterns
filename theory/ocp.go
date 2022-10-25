package theory

import "fmt"

/*
 *non ocp
 *when we need to add the new logic for other business, we need to modify the old code (Banker Class)
 */

// Banker ...
type Banker struct{}

// Deposit ...
func (b *Banker) Deposit() {
	fmt.Println("do the deposit business")
}

// Transfer ...
func (b *Banker) Transfer() {
	fmt.Println("do the transfer business")
}

// Payment ...
func (b *Banker) Payment() {
	fmt.Println("do the payment business")
}

// OtherBusiness for the new business logic
func (b *Banker) OtherBusiness() {
	fmt.Println("do the new business")
}

func logicNonOcp() {
	b := new(Banker)
	b.Deposit()
	b.Transfer()
	b.Payment()

	// new
	b.OtherBusiness()
}

/*
 *meet the ocp with interface(abstracts)
 */

// IBanker ...
type IBanker interface {
	DoBusiness()
}

// DepositBanker ...
type DepositBanker struct{}

// DoBusiness ...
func (db *DepositBanker) DoBusiness() {
	fmt.Println("do the deposit business")
}

// TransferBanker ...
type TransferBanker struct{}

// DoBusiness ...
func (db *TransferBanker) DoBusiness() {
	fmt.Println("do the transfer business")
}

// PaymentBanker ...
type PaymentBanker struct{}

// DoBusiness ...
func (db *PaymentBanker) DoBusiness() {
	fmt.Println("do the payment business")
}

// OtherBanker ...
type OtherBanker struct{}

// DoBusiness ...
func (ob *OtherBanker) DoBusiness() {
	fmt.Println("do the other business")
}

func logicOcp() {
	db := new(DepositBanker)
	db.DoBusiness()

	tb := new(TransferBanker)
	tb.DoBusiness()

	pb := new(PaymentBanker)
	pb.DoBusiness()

	// new logic
	ob := new(OtherBanker)
	ob.DoBusiness()
}

// for more abstract design

// Bank ...
func Bank(banker IBanker) {
	banker.DoBusiness()
}

func logic() {
	Bank(new(DepositBanker))
	Bank(new(TransferBanker))
	Bank(new(PaymentBanker))

	// new logic
	Bank(new(OtherBanker))
}
