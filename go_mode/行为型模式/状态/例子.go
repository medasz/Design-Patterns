package main

/*
type State interface {
	doAction(context *Context)
	toString()string
}

type StartState struct {
}
func (s *StartState)doAction(context *Context){
	println("Player is in start state")
	context.setState(s)
}
func (s *StartState)toString()string{
	return "Start State"
}

type StopState struct {
}
func (s *StopState)doAction(context *Context){
	println("Player is in stop state")
	context.setState(s)
}
func (s *StopState)toString()string{
	return "Stop State"
}


type Context struct {
	state State
}
func (c *Context)Context(){
	c.state=nil
}
func (c *Context)setState(state State){
	c.state=state
}
func (c *Context)getState()State{
	return c.state
}

func main(){
	context:=Context{}
	context.Context()

	startState:=StartState{}
	startState.doAction(&context)
	fmt.Println(context.getState().toString())

	stopState:=StopState{}
	stopState.doAction(&context)
	println(context.getState().toString())
}
*/