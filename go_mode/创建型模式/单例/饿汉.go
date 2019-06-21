package main


type example2 struct {
	name string
}
var instance2 *example2

func init(){
	instance2=new(example2)
	instance2.name="初始化单例模式"
}

func GetInstance2() *example2{
	return instance2
}
func main(){
	s:=GetInstance2()
	println(s.name)
}