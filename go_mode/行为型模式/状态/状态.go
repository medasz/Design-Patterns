package main

import "fmt"

type Week interface {
	Today()
	Next(*DayContext)
}
type Sunday struct {
}
func (*Sunday)Today(){
	fmt.Printf("Sunday\n")
}
func (*Sunday)Next(ctx *DayContext){
	ctx.today=&Monday{}
}

type Monday struct {
}
func (*Monday)Today(){
	fmt.Printf("Monday\n")
}
func (*Monday)Next(ctx *DayContext){
	ctx.today=&Tuesday{}
}

type Tuesday struct {
}
func (*Tuesday)Today(){
	fmt.Println("Tuesday")
}
func (*Tuesday)Next(ctx *DayContext){
	ctx.today=&Wednesday{}
}

type Wednesday struct {
}
func (*Wednesday)Today(){
	fmt.Println("Wednesday")
}
func (*Wednesday)Next(ctx *DayContext){
	ctx.today=&Thursday{}
}

type Thursday struct {
}
func (*Thursday)Today(){
	fmt.Println("Thursday")
}
func (*Thursday)Next(ctx *DayContext){
	ctx.today=&Friday{}
}

type Friday struct {
}
func (*Friday)Today(){
	fmt.Println("Friday")
}
func (*Friday)Next(ctx *DayContext){
	ctx.today=&Saturday{}
}

type Saturday struct {
}
func (*Saturday)Today(){
	fmt.Println("Saturday")
}
func (*Saturday)Next(ctx *DayContext){
	ctx.today=&Sunday{}
}


type DayContext struct {
	today Week
}
func NewDayContext() *DayContext{
	return &DayContext{
		today:&Sunday{},
	}
}
func (d *DayContext)Today(){
	d.today.Today()
}
func (d *DayContext)Next(){
	d.today.Next(d)
}
func main(){
	ctx:=NewDayContext()
	todayAndNext:= func() {
		ctx.Today()
		ctx.Next()
	}
	for i:=0;i<8;i++{
		todayAndNext()
	}
}