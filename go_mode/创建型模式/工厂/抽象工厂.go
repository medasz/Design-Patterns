package main


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
	CreateProductFirst() ProductInterface
	CreateProductSecond() ProductInterface
}
type Factory1 struct {
}
func (f Factory1)CreateProductFirst()ProductInterface{
	return 	&Product1{}
}
func (f Factory1)CreateProductSecond()ProductInterface{
	return &Product2{}
}


type Factory2 struct {
}
func (f Factory2)CreateProductFirst()ProductInterface{
	return 	&ProductA{}
}
func (f Factory2)CreateProductSecond()ProductInterface{
	return &ProductB{}
}

func main(){
	f1:=Factory1{}
	pi:=f1.CreateProductFirst()
	pi.Intro()
	pi=f1.CreateProductSecond()
	pi.Intro()

	f2:=Factory2{}
	pi=f2.CreateProductFirst()
	pi.Intro()
	pi=f2.CreateProductSecond()
	pi.Intro()
}