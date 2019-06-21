package main

type example struct {
	name string
}
var instance *example

func GetExample() *example{
	//存在线程安全问题，高并发时有可能创建多个对象
	if instance==nil{
		instance=new(example)
	}
	return instance
}

func main(){
	s:=GetExample()
	s.name="第一次赋值单例模式"
	println(s.name)

	s2:=GetExample()
	println(s2.name)
}
