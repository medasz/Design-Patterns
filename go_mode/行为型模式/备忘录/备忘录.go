package main

//备忘录
type Memento struct {
	state string//保存的状态
}
func (m *Memento)SetState(s string){
	m.state=s
}
func (m *Memento)GetState()string{
	return m.state
}


//发起人结构体
type Originator struct {
	state string//要保存的状态
}
func (o *Originator)SetState(s string){
	o.state=s
}
func (o *Originator)GetState()string{
	return o.state
}
//规定保存的状态范围
func (o *Originator)CreateMemento()*Memento{
	return &Memento{
		state:o.state,
	}
}


//负责人结构体
type Caretaker struct {
	memento *Memento
}
func (c *Caretaker)GetMemento()*Memento{
	return c.memento
}
func (c *Caretaker)SetMemento(m *Memento){
	c.memento=m
}

func main(){
	//创建一个发起人并设置初始状态
	o:=&Originator{state:"hello"}
	println("当前状态:",o.GetState())
	//保存当前状态
	//创建负责人设置
	c:=new(Caretaker)
	c.SetMemento(o.CreateMemento())

	o.SetState("world")
	println("更改当前状态:",o.GetState())

	//恢复备忘
	o.SetState(c.GetMemento().GetState())
	println("恢复后状态",o.GetState())
}
