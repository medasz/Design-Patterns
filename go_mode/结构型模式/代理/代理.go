package main

import "reflect"

type Image interface {
	display()
}

type RealImage struct {
	fileName string
}
func (r *RealImage)RealImage(fileName string){
	r.fileName=fileName
	r.loadFromDisk(fileName)
}
func (r RealImage)display(){
	println("Displaying",r.fileName)
}
func (r RealImage)loadFromDisk(fileName string){
	println("Loading",fileName)
}

type ProxyImage struct {
	realImage RealImage
	fileName string
}
func (p *ProxyImage)ProxyImage(fileName string){
	p.fileName=fileName
}
func (p *ProxyImage)display(){
	if reflect.DeepEqual(p.realImage,RealImage{}){
		p.realImage=RealImage{}
		p.realImage.RealImage(p.fileName)
	}
	p.realImage.display()
}

func main(){
	ima:=ProxyImage{}
	ima.ProxyImage("test.jpg")
	ima.display()
	println()
	ima.display()
}