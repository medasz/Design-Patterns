package main

type StartStop interface {
	startup()
	shutdown()
}

type CPU struct {
}
func (c CPU)startup(){
	println("cpu startup!")
}
func (c CPU)shutdown(){
	println("cpu shutdown!")
}

type Memory struct {
}
func (m Memory)startup(){
	println("memory startup!")
}
func (m Memory)shutdown(){
	println("menory shutdown!")
}

type Disk struct {
}
func (d Disk)startup(){
	println("disk startup!")
}
func (d Disk)shutdown(){
	println("disk shutdown!")
}


type Computer struct {
	cpu StartStop
	memory StartStop
	disk StartStop
}
func (c *Computer)NewComputer(){
	c.cpu=new(CPU)
	c.memory=new(Memory)
	c.disk=new(Disk)
}
func (c Computer)startup(){
	println("start the computer!")
	c.cpu.startup()
	c.memory.startup()
	c.disk.startup()
	println("start computer finished!")
}

func (c Computer)shutdown(){
	println("begin to close the computer!")
	c.cpu.shutdown()
	c.memory.shutdown()
	c.disk.shutdown()
	println("computer close!")
}

func main(){
	c:=Computer{}
	c.NewComputer()
	c.startup()
	c.shutdown()
}