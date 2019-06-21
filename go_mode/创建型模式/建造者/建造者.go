package main

type Car struct {
	Wheels  string
	Chassis string
	Seat    string
}

func (c *Car) Show() {
	println(c.Wheels)
	println(c.Chassis)
	println(c.Seat)
}

type CarBuilder struct {
	Car *Car
}

func (cb *CarBuilder) GetResult() interface{} {
	return cb.Car
}
func (cb *CarBuilder) NewProduct() {
	cb.Car = new(Car)
}
func (cb *CarBuilder) BuildWheels() {
	cb.Car.Wheels = "build wheels"
}
func (cb *CarBuilder) BuildChassis() {
	cb.Car.Chassis = "build chassis"
}
func (cb *CarBuilder) BuildSeat() {
	cb.Car.Seat = "build seat"
}

type Builder interface {
	NewProduct()
	BuildWheels()
	BuildChassis()
	BuildSeat()
	GetResult()interface{}
}

type Director struct {
	builder Builder
}
func (d *Director)SetBuilder(builder Builder){
	d.builder=builder
}
func (d *Director)Generate()*Car{
	d.builder.NewProduct()
	d.builder.BuildChassis()
	d.builder.BuildSeat()
	d.builder.BuildWheels()
	return d.builder.GetResult().(*Car)
}

func main() {
	//创建指挥者
	director := new(Director)
	//创建建造者
	builder:=new(CarBuilder)
	director.SetBuilder(builder)
	car:=director.Generate()
	car.Show()
}
