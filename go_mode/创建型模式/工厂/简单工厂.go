package main
/*
type Factory struct {
}

func (f Factory)Generate(name string)Product{
	switch name {
	case "product1":
		return Product1{}
	case "product2":
		return Product2{}
	default:
		return nil
	}
}


type Product interface {
	create()
}

type Product1 struct {
}
func (p1 Product1)create(){
	println("product1")
}

type Product2 struct {
}
func (p2 Product2)create(){
	println("product2")
}

func main(){
	f:=Factory{}
	api1:=f.Generate("product1")
	api1.create()
	api2:=f.Generate("product2")
	api2.create()
}
*/