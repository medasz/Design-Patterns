package main

import "fmt"

type MotherBoard struct {
}
func (*MotherBoard)Start(){
	fmt.Println("system starting")
}
func (*MotherBoard)Reboot(){
	fmt.Println("system rebooting")
}


type Command interface {
	Execute()
}
type StartCommand struct {
	mb *MotherBoard
}
func NewStartCommand(mb *MotherBoard)*StartCommand{
	return &StartCommand{
		mb:mb,
	}
}
func (c *StartCommand)Execute(){
	c.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}
func NewRebootCommand(mb *MotherBoard)*RebootCommand{
	return &RebootCommand{
		mb:mb,
	}
}
func (r *RebootCommand)Execute(){
	r.mb.Reboot()
}

type Box struct {
	button1 Command
	button2 Command
}
func NewBox(button1,button2 Command)*Box{
	return &Box{
		button1:button1,
		button2:button2,
	}
}
func (b *Box)PressButtion1(){
	b.button1.Execute()
}
func (b *Box)PressButtion2(){
	b.button2.Execute()
}

func main(){
	mb:=&MotherBoard{}
	startCommand:=NewStartCommand(mb)
	rebootCommand:=NewRebootCommand(mb)

	box1:=NewBox(startCommand,rebootCommand)
	box1.PressButtion1()
	box1.PressButtion2()

	box2:=NewBox(rebootCommand,startCommand)
	box2.PressButtion1()
	box2.PressButtion2()

}