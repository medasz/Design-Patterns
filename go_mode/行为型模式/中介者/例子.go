package main

import (
	"time"
)

type ChatRoom struct {
}
func (c ChatRoom)showMessage(user User,message string){
	println(time.Now().String()+"["+user.getName()+"]:"+message)
}


type User struct {
	chat *ChatRoom
	name string
}
func (u User)getName()string{
	return u.name
}
func (u User)setName(name string){
	u.name=name
}
func (u User)sendMessage(message string){
	ChatRoom.showMessage(ChatRoom{},u,message)
}
func NewUser(name string)*User{
	return &User{
		name:name,
	}
}

func main(){
	robert:=NewUser("Robert")
	john:=NewUser("John")
	robert.chat=&ChatRoom{}
	john.chat=&ChatRoom{}
	robert.sendMessage("Hi!John!")
	john.sendMessage("Hello!Robert!")
}