package main

type Base struct {
}

func (b Base) Print() {
	println("这是print方法")
}
func (b Base) Echo() {
}

type Son struct {
	Base
}

func (s Son) Echo() {
	println("这是Echo方法")
}

func main() {
	s := new(Son)
	s.Print()
	s.Echo()
}
