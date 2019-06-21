package main
/*
import "fmt"

type IObserver interface {
	Notify()
}
type Observer struct {
}
func (o *Observer)Notify(){
	fmt.Println("已经触发了观察者")
}

type ISubject interface {
	AddObservers(observers ...IObserver) //添加观察者
	NotifyObservers()                    //通知观察者
}
type Subject struct {
	observers []IObserver
}
func (s *Subject)AddObservers(observers ...IObserver){
	s.observers=append(s.observers,observers...)
}
func (s *Subject)NotifyObservers(){
	for k:=range s.observers{
		s.observers[k].Notify()
	}
}
*/

func main(){
	//创建被观察者
	s:=new(Subject)
	//创建观察者
	o:=new(Observer)
	//为主题添加观察者
	s.AddObservers(o)
	s.NotifyObservers()
}
