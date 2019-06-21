package main

/*
type ProductInterface interface {
	Intro()
}
type Product1 struct {
}
func (p Product1)Intro(){
	println("product1")
}
type Product2 struct {
}
func (p Product2)Intro(){
	println("product2")
}
type ProductA struct {
}
func (p ProductA)Intro(){
	println("productA")
}
type ProductB struct {
}
func (p ProductB)Intro(){
	println("productB")
}



type FactoryInterface interface {
	CreateProduct(t string)ProductInterface
}
type Factory1 struct {
}
func (f Factory1)CreateProduct(t string)ProductInterface{
	switch t {
	case "product1":
		return &Product1{}
	case "product2":
		return &Product2{}
	default:
		return nil
	}
}

type Factory2 struct {
}
func (f Factory2)CreateProduct(t string)ProductInterface{
	switch t {
	case "productA":
		return &ProductA{}
	case "productB":
		return &ProductB{}
	default:
		return nil
	}
}


func main(){
	var pi ProductInterface
	f1:=Factory1{}
	pi=f1.CreateProduct("product1")
	pi.Intro()
	pi=f1.CreateProduct("product2")
	pi.Intro()
	f2:=Factory2{}
	pi=f2.CreateProduct("productA")
	pi.Intro()
	pi=f2.CreateProduct("productB")
	pi.Intro()
}
*/