package main

import "fmt"

type IStrategy interface {
	SortList()
}
type BubbleSortStrategy struct {
}
func (b BubbleSortStrategy)SortList(){
	fmt.Println("这是冒泡排序")
}

type MergeSortStrategy struct {
}
func (m MergeSortStrategy)SortList(){
	fmt.Println("这是归并排序")
}


type Context struct {
	Strategy IStrategy
}
func (c Context)Exec(){
	c.Strategy.SortList()
}

func main(){
	var ctx Context
	fmt.Println("====使用冒泡排序算法=====")
	ctx=Context{Strategy:BubbleSortStrategy{}}
	ctx.Exec()

	fmt.Println("====使用归并排序算法=====")
	ctx=Context{Strategy:MergeSortStrategy{}}
	ctx.Exec()
}
