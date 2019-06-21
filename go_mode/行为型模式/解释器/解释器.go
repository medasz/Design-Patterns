package main

import (
	"strconv"
	"strings"
)

type Node interface {
	Interpret() int
}
type ValNode struct {
	val int
}

func (n *ValNode) Interpret() int {
	return n.val
}

type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

type MinNode struct {
	left, right Node
}

func (n *MinNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

type Parser struct {
	exp   []string
	index int
	prev  Node
}
//1 + 2 + 3
func (p *Parser)Parse(exp string){
	p.exp=strings.Split(exp," ")
	for{
		if p.index>=len(p.exp){
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev=p.newAddNode()
		case "-":
			p.prev=p.newMinNode()
		default:
			p.prev=p.newValNode()
		}
	}
}
func (p *Parser)newAddNode()Node{
	p.index++
	return &AddNode{
		left:p.prev,
		right:p.newValNode(),
	}
}
func (p *Parser)newMinNode()Node{
	p.index++
	return &MinNode{
		left:p.prev,
		right:p.newValNode(),
	}
}
func (p *Parser)newValNode()Node{
	v,_:=strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{
		val:v,
	}
}
func (p *Parser)Result()Node{
	return p.prev
}

func main(){
	p:=&Parser{}
	p.Parse("1 + 3 + 4 + 9")
	res:=p.Result().Interpret()
	println(res)
}