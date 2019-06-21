package main

import (
	"sync"
	"sync/atomic"
)

type example struct {
	name string
}
var instance *example


var mux sync.Mutex
var initialized uint32
func GetExample() *example{
	//一次判断即可返回
	if atomic.LoadUint32(&initialized)==1{
		return instance
	}
	mux.Lock()
	defer mux.Unlock()
	if initialized==0{
		instance=new(example)
		atomic.StoreUint32(&initialized,1)
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