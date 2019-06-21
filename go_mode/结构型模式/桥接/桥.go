package main

import "strconv"

//创建桥接实现接口
type DrawAPI interface {
	drawCircle(radius int,x int,y int)
}

//创建实现了DrawAPI接口的实体桥接实现类
type RedCircle struct {
}
func (r RedCircle)drawCircle(radius int,x int,y int){
	println("Drawing Circle[ color: red,radius:"+strconv.Itoa(radius)+",x:"+strconv.Itoa(x)+",y:"+strconv.Itoa(y)+"]")
}
type GreenCircle struct {
}
func (g GreenCircle)drawCircle(radius,x,y int){
	println("Drawing Circle[ color: green,radius:"+strconv.Itoa(radius)+",x:"+strconv.Itoa(x)+",y:"+strconv.Itoa(y)+"]")
}

//使用DrawAPI接口创建抽象类Shape
type Shape struct {
	drawAPI DrawAPI
}
func (s *Shape)Shape(drawAPI DrawAPI){
	s.drawAPI=drawAPI
}
func (s *Shape)draw(){}

//创建实现了Shape接口的实体类
type Circle struct {
	s Shape
	x,y,radius int
}
func (c *Circle)Circle(x,y,radius int,drawAPI DrawAPI){
	c.x=x
	c.y=y
	c.radius=radius
	c.s.drawAPI=drawAPI
}
func (c Circle)draw(){
	c.s.drawAPI.drawCircle(c.radius,c.x,c.y)
}


func main(){
	redCircle:=Circle{}
	redCircle.Circle(100,100,10,RedCircle{})

	greenCircle:=Circle{}
	greenCircle.Circle(100,100,10,GreenCircle{})

	redCircle.draw()
	greenCircle.draw()
}