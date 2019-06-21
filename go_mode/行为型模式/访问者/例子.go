package main

import (
	"fmt"
)

type Customer interface {
	Accept(Visitor)
}

type EnterpriseCustomer struct {
	name string
}
func NewEnterpriseCustomer(name string)*EnterpriseCustomer{
	return &EnterpriseCustomer{name}
}
func (c *EnterpriseCustomer)Accept(visitor Visitor){
	visitor.Visit(c)
}

type IndividualCustomer struct {
	name string
}
func NewIndividualCustomer(name string)*IndividualCustomer{
	return &IndividualCustomer{
		name:name,
	}
}
func (c *IndividualCustomer)Accept(visitor Visitor){
	visitor.Visit(c)
}


type CustomerCol struct {
	customers []Customer
}
func (c *CustomerCol)Add(customer Customer){
	c.customers=append(c.customers,customer)
}
func (c *CustomerCol)Accept(visitor Visitor){
	for _,customer:=range c.customers{
		customer.Accept(visitor)
	}
}






type Visitor interface {
	Visit(Customer)
}

type ServieRequestVisitor struct {
}
func (*ServieRequestVisitor)Visit(customer Customer){
	switch c:=customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n",c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n",c.name)
	}
}

type AnalysisVisitor struct {
}
func (*AnalysisVisitor)Visit(customer Customer){
	switch c:=customer.(type){
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n",c.name)
	case *IndividualCustomer:
		fmt.Printf("analysis individual customer %s\n",c.name)
	}
}


func main(){
	c:=&CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Accept(&ServieRequestVisitor{})

	a:=&CustomerCol{}
	a.Add(NewEnterpriseCustomer("A company"))
	a.Add(NewIndividualCustomer("bob"))
	a.Add(NewEnterpriseCustomer("B company"))
	a.Accept(&AnalysisVisitor{})
}
