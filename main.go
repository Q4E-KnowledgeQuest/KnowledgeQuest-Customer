package main

import (
	"flag"
	"main/components/courses"
	"main/server"
	"github.com/jasonlvhit/gocron"
)

func main() {
	portPtr := flag.Int("port", 3000, "Port to listen on")
	flag.Parse()
	courses.DownloadCourses()
	go Schedule()
	server.Start(*portPtr)
}

func Schedule() {
	gocron.Every(1).Hour().Do(courses.DownloadCourses)
	<-gocron.Start()
}