package main

import "fmt"

type Downloader interface {
	Download(uri string)
}
type template struct {
	implement
	uri string
}

func newTemplate(impl implement) *template {
	return &template{
		impl,
		"",
	}
}
func (t *template) Download(uri string) {
	t.uri = uri
	println("prepare downloading")
	t.implement.download()
	t.implement.save()
	println("finish downloading")
}
func (t *template) save() {
	println("default save")
}

type implement interface {
	download()
	save()
}
type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}
func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}
func (d *HTTPDownloader) save() {
	println("http save")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}
func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}


func main(){
	var downloader Downloader=NewHTTPDownloader()
	downloader.Download("http://example.com/abc.zip")

	var downloaderftp Downloader=NewFTPDownloader()
	downloaderftp.Download("ftp://example.com/abc.zip")
}