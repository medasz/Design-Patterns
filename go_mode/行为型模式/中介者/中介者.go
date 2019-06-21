package main

import (
	"fmt"
	"log"
	"strings"
)

type CDDriver struct {
	Data string
}
func (c *CDDriver)ReadData(){
	c.Data="music,image"
	fmt.Printf("CDDriver: reading data %s\n",c.Data)
	GetMediatorInstance().changed(c)
}

type CPU struct {
	Video string
	Sound string
}
func (c *CPU)Process(data string){
	sp:=strings.Split(data,",")
	c.Sound=sp[0]
	c.Video=sp[1]

	fmt.Printf("CPU: split data with Sound %s, Video %s\n",c.Sound,c.Video)
	GetMediatorInstance().changed(c)
}

type VideoCard struct {
	Data string
}
func (v *VideoCard)Display(data string){
	v.Data=data
	fmt.Printf("VideoCard: display %s\n",v.Data)
	GetMediatorInstance().changed(v)
}

type SoundCard struct {
	Data string
}
func (s *SoundCard)Play(data string){
	s.Data=data
	fmt.Printf("SoundCard: play %s\n",s.Data)
	GetMediatorInstance().changed(s)
}

var mediator *Mediator
type Mediator struct {
	CD *CDDriver
	CPU *CPU
	Video *VideoCard
	Sound *SoundCard
}
func GetMediatorInstance() *Mediator{
	if mediator==nil{
		mediator=&Mediator{}
	}
	return mediator
}
func (m *Mediator)changed(i interface{}){
	switch inst:=i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	}
}

func main(){
	mediator :=GetMediatorInstance()
	mediator.CD=&CDDriver{}
	mediator.CPU=&CPU{}
	mediator.Video=&VideoCard{}
	mediator.Sound=&SoundCard{}

	mediator.CD.ReadData()
	if mediator.CD.Data!="music,image"{
		log.Fatal("CD unexcept data %s",mediator.CD.Data)
	}

	if mediator.CPU.Sound!="music"{
		log.Fatal("CPU unexcept sound data %s",mediator.CPU.Sound)
	}

	if mediator.CPU.Video!="image"{
		log.Fatal("CPU unexcept sound data %s",mediator.CPU.Video)
	}

	if mediator.Video.Data!="image"{
		log.Fatal("CPU unexcept sound data %s",mediator.Video.Data)
	}

	if mediator.Sound.Data!="music"{
		log.Fatal("CPU unexcept sound data %s",mediator.Sound.Data)
	}
}