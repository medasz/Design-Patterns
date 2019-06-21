package main

type IVisitor interface {
	Visit()
}

type ProductionVisitor struct {
}
func (v ProductionVisitor)Visit(){
	println("这是生产环境")
}

type TestingVisitor struct {
}
func (t TestingVisitor)Visit(){
	println("这是测试环境")
}

type IElement interface {
	Accept(visitor IVisitor)
}
type Element struct {
}
func (el Element)Accept(visitor IVisitor){
	visitor.Visit()
}

type EnvExample struct {
	Element
}
func (e EnvExample)Print(visitor IVisitor){
	e.Element.Accept(visitor)
}

func main(){
	e:=new(Element)
	e.Accept(new(ProductionVisitor))
	e.Accept(new(TestingVisitor))

	m:=new(EnvExample)
	m.Print(new(ProductionVisitor))
	m.Print(new(TestingVisitor))
}