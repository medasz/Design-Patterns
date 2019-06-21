package main
/*
import "fmt"

//容器接口
type IAggregate interface {
	Iterator() IIterator
}

//实现具体容器
type Aggregate struct {
	container []int //容器中装载int型容器
}

func (a *Aggregate) Iterator() IIterator {
	i := new(Iterator)
	i.aggregate=a
	return i
}

//迭代器接口
type IIterator interface {
	HasNext() bool
	Current() int
	Next() bool
}

//具体迭代器
type Iterator struct {
	cursor    int        //当前游标
	aggregate *Aggregate //对应的容器指针
}
//判断是否迭代到最后，如果没有，则返回true
func (i *Iterator)HasNext()bool{
	if i.cursor+1<len(i.aggregate.container){
		return true
	}
	return false
}
//获取当前迭代元素（从容器中取出当前游标对应的元素）
func (i *Iterator)Current()int{
	return i.aggregate.container[i.cursor]
}
//将游标指向下一个元素
func (i *Iterator)Next()bool{
	if i.HasNext(){
		i.cursor++
		return true
	}
	return false
}

func main(){
	//创建容器，并放入初始化数据
	c:=&Aggregate{container:[]int{1,2,3,4}}
	//获取迭代器
	iterator:=c.Iterator()
	for{
		//打印当前数据
		fmt.Println(iterator.Current())
		if iterator.HasNext(){
			iterator.Next()
		}else{
			break
		}
	}
}
*/
