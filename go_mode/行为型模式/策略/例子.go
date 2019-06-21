package main

import "fmt"

type PaymentContext struct {
	Name,CardID string
	Money int
	payment PaymentStrategy
}
func NewPaymentContext(name ,cardif string,money int,payment PaymentStrategy)*PaymentContext{
	return &PaymentContext{
		Name:name,
		CardID:cardif,
		Money:money,
		payment:payment,
	}
}
func (p *PaymentContext)Pay(){
	p.payment.Pay(p)
}






type PaymentStrategy interface {
	Pay(*PaymentContext)
}


type Cash struct {
}
func (*Cash)Pay(ctx *PaymentContext){
	fmt.Printf("Pay $%d to %s by bank account %s\n",ctx.Money,ctx.Name,ctx.CardID)
}

type Bank struct {
}
func (*Bank)Pay(ctx *PaymentContext){
	fmt.Printf("Pay $%d to %s by bank account %s\n",ctx.Money,ctx.Name,ctx.CardID)
}


func main(){
	ctx:=NewPaymentContext("Ada","",123,&Cash{})
	ctx.Pay()

	cyx:=NewPaymentContext("Bob","0002",888,&Bank{})
	cyx.Pay()
}
