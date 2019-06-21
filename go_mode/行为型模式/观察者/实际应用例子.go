package main

import "fmt"

//创建被观察者
type Subject struct {
	observer []Observer
	context string
}
func NewSubject()*Subject{
	return &Subject{
		observer:make([]Observer,0),
	}
}
func (s *Subject)Attach(o Observer){
	s.observer=append(s.observer,o)
}
func (s *Subject)notify(){
	for _,o:=range s.observer{
		o.Update(s)
	}
}
func (s *Subject)UpdateContext(context string){
	s.context=context
	s.notify()
}

//创建观察者
type Observer interface {
	Update(*Subject)
}
type Reader struct {
	name string
}
func NewReader(name string)*Reader{
	return &Reader{
		name:name,
	}
}
func (r *Reader)Update(s *Subject){
	fmt.Printf("%s receive %s\n",r.name,s.context)
}

func main(){
	subject:=NewSubject()
	reader1:=NewReader("reader1")
	reader2:=NewReader("reader2")
	reader3:=NewReader("reader3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")
}