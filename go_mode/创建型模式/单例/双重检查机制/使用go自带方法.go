package main

import "sync"

type example3 struct {
	name string
}
var instance3 *example3
var once sync.Once

func GetInstance3()*example3{
	once.Do(func() {
		instance3=new(example3)
		instance3.name="第一次赋值单例"
	})
	return instance3
}

func main(){
	e1:=GetInstance3()
	println(e1.name)
	e2:=GetInstance3()
	println(e2.name)
}