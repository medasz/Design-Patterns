package main

import (
	"github.com/levigross/grequests"
	"log"
)

//
//import (
//	"github.com/levigross/grequests"
//	"log"
//)
//
//type sendRequest struct {
//	url string
//	observer []MyObserver
//}
//func NewRequest(url string)*sendRequest{
//	return &sendRequest{
//		url:url,
//	}
//}
//func (s *sendRequest)Send(){
//	res,err:=grequests.Get(s.url,nil)
//	if err!=nil{
//		log.Fatal(err)
//	}
//	println(res.String())
//}
//func (s *sendRequest)Attach(observer MyObserver){
//	s.observer=append(s.observer,observer)
//}
//func (s *sendRequest)noitfy(){
//
//}
//
//
//type MyObserver interface {
//	notify()
//}
func main(){
	res,err:=grequests.Get("https://cn-proxy.com/",nil)
	if err!=nil{
		log.Fatal(err)
	}
	println(res.String())
}