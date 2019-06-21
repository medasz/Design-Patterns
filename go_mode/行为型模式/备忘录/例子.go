package main
/*
import (
	"fmt"
)


type gameMemento struct {
	hp,mp int
}
type Memento interface {}


type Game struct {
	hp,mp int
}

func (g *Game)Play(mpDelta,hpDelta int){
	g.hp+=hpDelta
	g.mp+=mpDelta
}
func (g *Game)Save()Memento{
	return &gameMemento{
		hp:g.hp,
		mp:g.mp,
	}
}
func (g *Game)Load(m Memento){
	gm:=m.(*gameMemento)
	g.mp=gm.mp
	g.hp=gm.hp
}
func (g *Game)Status(){
	fmt.Printf("Current HP:%d, MP:%d\n",g.hp,g.mp)
}

func main(){
	game:=&Game{
		hp:10,
		mp:10,
	}
	game.Status()
	progress:=game.Save()

	game.Play(-2,-3)
	game.Status()

	game.Load(progress)
	game.Status()
}

 */