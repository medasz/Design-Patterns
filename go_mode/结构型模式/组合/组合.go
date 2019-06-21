package main

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}

type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}
func (c *component) SetParent(parent Component) {
	c.parent = parent
}
func (c *component) Name() string {
	return c.name
}
func (c *component) SetName(name string) {
	c.name = name
}
func (c *component) AddChild(Component) {
}
func (c *component) Print() {
}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}
func (c *Leaf) Print(pre string) {
	println(pre + "-" + c.Name())
}

type Composite struct {
	component
	childs []Component
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}
func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childs = append(c.childs, child)
}
func (c *Composite) Print(pre string) {
	println(pre+"+"+c.Name())
	pre+=" "
	for _,comp:=range c.childs{
		comp.Print(pre)
	}
}

func main(){
	root:=NewComponent(CompositeNode,"root")
	c1:=NewComponent(CompositeNode,"c1")
	c2:=NewComponent(CompositeNode,"c2")
	c3:=NewComponent(CompositeNode,"c3")

	l1:=NewComponent(LeafNode,"l1")
	l2:=NewComponent(LeafNode,"l2")
	l3:=NewComponent(LeafNode,"l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Print("")
}