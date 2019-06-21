package main

import (
	"fmt"
	"log"
	)

type ImageFlyweight struct {
	data string
}
func NewImageFlyweight(filename string)*ImageFlyweight{
	//load image file
	data:=fmt.Sprintf("image data %s",filename)
	return &ImageFlyweight{data:data}
}
func (i *ImageFlyweight)Data()string{
	return i.data
}

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}
var imageFactory *ImageFlyweightFactory
func GetImageFlyweightFactory()*ImageFlyweightFactory{
	if imageFactory==nil{
		imageFactory=&ImageFlyweightFactory{
			maps:make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}
func (f *ImageFlyweightFactory)Get(filename string)*ImageFlyweight{
	image:=f.maps[filename]
	if image==nil{
		image=NewImageFlyweight(filename)
		f.maps[filename]=image
	}
	return image
}


type ImageViewer struct {
	*ImageFlyweight
}
func NewImageViewer(filename string)*ImageViewer{
	image:=GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		image,
	}
}
func (i *ImageViewer)Display(){
	println("Display:",i.Data())
}

func main(){
	viewer:=NewImageViewer("image1.png")
	viewer.Display()

	viewer1:=NewImageViewer("image1.png")
	viewer2:=NewImageViewer("image1.png")
	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		log.Fatal("asd")
	}
}