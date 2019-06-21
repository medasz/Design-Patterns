package main


type BrushInterface interface {
	draw()
}
type BrushShape struct {
	Shape string
}
func NewBrushShape(Shape string)BrushInterface{
	return &BrushShape{Shape:Shape}
}
func (b BrushShape)draw(){
	println(b.Shape+"号笔,")
}


type BrushColorInterface interface {
	drawColor()
}
type Color struct {
	color string
	brushInterface BrushInterface
}
func NewColor(color string,brushInterface BrushInterface)*Color{
	return &Color{color:color,brushInterface:brushInterface}
}
func (c Color)drawColor(){
	c.brushInterface.draw()
	println("粘"+c.color+"色")
	println()
}


func main(){
	big:=NewBrushShape("大")
	mid:=NewBrushShape("中")
	sml:=NewBrushShape("小")

	c1:=NewColor("红",big)
	c2:=NewColor("红",mid)
	c3:=NewColor("红",sml)
	c1.drawColor()
	c2.drawColor()
	c3.drawColor()

	c4:=NewColor("蓝",big)
	c5:=NewColor("蓝",mid)
	c6:=NewColor("蓝",sml)
	c4.drawColor()
	c5.drawColor()
	c6.drawColor()
}