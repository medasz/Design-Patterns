package main

type Component interface {
	Calc() int
}
type ConcreteComponent struct {
}
func (*ConcreteComponent)Calc()int{
	return 0
}


type MulDecorator struct {
	Component
	num int
}
func WarpMulDecorator(c Component,num int)Component{
	return &MulDecorator{
		c,
		num,
	}
}
func (d *MulDecorator)Calc()int{
	return d.Component.Calc()*d.num
}

type AddDecorator struct {
	Component
	num int
}
func WarpAdddecorator(c Component,num int)Component{
	return &AddDecorator{
		c,
		num,
	}
}
func (d *AddDecorator)Calc()int{
	return d.Component.Calc()+d.num
}

func main(){
	var c Component=&ConcreteComponent{}
	c=WarpAdddecorator(c,10)
	c=WarpMulDecorator(c,8)
	res:=c.Calc()
	println("res ",res)
}