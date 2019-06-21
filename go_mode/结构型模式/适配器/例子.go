package main

import "log"

//模拟计算机读取SD卡
//sd卡接口
type SDCard interface {
	readSD()string
	writeSD(msg string)int
}
//sd卡
type SDCardImpl struct {
}
func (s SDCardImpl)readSD()string{
	msg:="sdcard read a msg:hello word SD"
	return msg
}
func (s SDCardImpl)writeSD(msg string)int{
	println("sd card write msg:"+msg)
	return 1
}

//计算机接口，读取sd卡
type Computer interface {
	readSD(sdCard SDCard)string
}
//计算机实例
type ThinkpadComputer struct {
}
func (t ThinkpadComputer)readSD(sdCard SDCard)string{
	if sdCard==nil{
		log.Fatal("sd card nil")
	}
	return sdCard.readSD()
}




//TF卡接口
type TFCard interface {
	readTF()string
	writeTF(msg string)int
}
//TF卡实例
type TFCardImpl struct {
}
func (tf TFCardImpl)readTF()string{
	msg:="tf card readmsg:hello word tf card"
	return msg
}
func (tf TFCardImpl)writeTF(msg string)int{
	println("tf card write a msg:"+msg)
	return 1
}


//创建SD适配TF
type SDAdapterTF struct {
	tfCard TFCard
}
func (sd *SDAdapterTF)SDAdapterTF(tfCard TFCard){
	sd.tfCard=tfCard
}
func (sd SDAdapterTF)readSD()string{
	println("adapter read tf card")
	return sd.tfCard.readTF()
}
func (sd SDAdapterTF)writeSD(msg string)int{
	println("adapter write tf card")
	return sd.tfCard.writeTF(msg)
}




func main(){
	computer:=new(ThinkpadComputer)
	sdCard:=new(SDCardImpl)
	println(computer.readSD(sdCard))
	println("=================================")
	tfCard:=new(TFCardImpl)
	tfCardAdapterSD:=SDAdapterTF{}
	tfCardAdapterSD.SDAdapterTF(tfCard)
	println(computer.readSD(tfCardAdapterSD))
}