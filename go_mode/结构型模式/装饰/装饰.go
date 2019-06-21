package main

type Shape interface {
	draw()
}

type Rectangle struct {
}
func (r Rectangle)draw(){
	println("Shape: Rectangle")
}

type Circle struct {
}
func (c Circle)draw(){
	println("Shape: Circle")
}


type ShapeDecorator struct {
	decoratedShape Shape
}
func (s *ShapeDecorator)ShapeDecorator(decoratedShape Shape){
	s.decoratedShape=decoratedShape
}
func (s ShapeDecorator)draw(){
	s.decoratedShape.draw()
}


type RedShapeDecorator struct {
	s ShapeDecorator
}
func (r *RedShapeDecorator)RedShapeDecorator(decoratedShape Shape){
	r.s.ShapeDecorator(decoratedShape)
}
func (r *RedShapeDecorator)setRedBorder(decoratedShape Shape){
	println("Border Color: Red")
}
func (r RedShapeDecorator)draw(){
	r.s.decoratedShape.draw()
	r.setRedBorder(r.s)
}

func main(){
	c:=Circle{}
	redCircle:=RedShapeDecorator{}
	redCircle.RedShapeDecorator(c)
	r:=Rectangle{}
	redRectangle:=RedShapeDecorator{}
	redRectangle.RedShapeDecorator(r)

	println("Circle with normal border")
	c.draw()

	println("Circle of red border")
	redCircle.draw()

	println("Rectangle of red border")
	redRectangle.draw()
}