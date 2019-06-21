package main

type Shape interface {
	draw()
}

type Rectangle struct {
}

func (r *Rectangle) draw() {
	println("Rectangle::draw()")
}

type Square struct {
}

func (s *Square) draw() {
	println("Square::draw()")
}

type Circle struct {
}

func (c *Circle) draw() {
	println("Circle::draw()")
}

type ShapeMaker struct {
	circle    Shape
	rectangle Shape
	square    Shape
}
func (s *ShapeMaker)NewSystem(){
	s.circle=new(Circle)
	s.rectangle=new(Rectangle)
	s.square=new(Square)
}
func (s *ShapeMaker)drawCircle(){
	s.circle.draw()
}
func (s *ShapeMaker)drawRectangle(){
	s.rectangle.draw()
}
func (s *ShapeMaker)drawSquare() {
	s.square.draw()
}


func main(){
	sm:=ShapeMaker{}
	sm.NewSystem()
	sm.drawCircle()
	sm.drawRectangle()
	sm.drawSquare()
}