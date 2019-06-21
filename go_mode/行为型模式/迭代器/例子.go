package main

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}
type Numbers struct {
	start,end int
}
func NewNumbers(start,end int)*Numbers{
	return &Numbers{
		start:start,
		end:end,
	}
}
func (n *Numbers)Iterator()Iterator{
	return &NumbersIterator{
		numbers:n,
		next:n.start,
	}
}

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}
type NumbersIterator struct {
	numbers *Numbers
	next int
}
func (i *NumbersIterator)First(){
	i.next=i.numbers.start
}
func (i *NumbersIterator)IsDone()bool{
	return i.next>i.numbers.end
}
func (i *NumbersIterator)Next()interface{}{
	if !i.IsDone(){
		next:=i.next
		i.next++
		return next
	}
	return nil
}
func IteratorPrint(i Iterator){
	for i.First();!i.IsDone();{
		c:=i.Next()
		fmt.Printf("%#v\n",c)
	}
}


func main(){
	var aggregate Aggregate
	aggregate=NewNumbers(1,10)
	IteratorPrint(aggregate.Iterator())
}
