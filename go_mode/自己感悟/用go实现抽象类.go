package main

type abstruct struct {
}
func (abstruct)eat(){
	println("abstract eat")
}
func (abstruct)speak(){
	println("abstract speak")
}



type Person struct {
	abstruct
}
func NewPerson(a abstruct)*Person{
	return &Person{
		a,
	}
}

func main(){
	a:=abstruct{}

	p:=Person{a}
	p.eat()
	p.speak()
}

