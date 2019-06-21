package main

import "fmt"

type Example struct {
	Content string
}

func (e *Example)Clone() *Example{
	res:=*e
	return &res
}

func main(){
	r1:=new(Example)
	r1.Content="example1"

	r2:=r1.Clone()
	r2.Content="example2"

	fmt.Println(r1)
	fmt.Println(r2)
}

