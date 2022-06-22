package main

import (
	"flag"
	"main/components/courses"
	"main/server"
	"github.com/jasonlvhit/gocron"
	"github.com/pterm/pterm"
)

func main() {
	portPtr := flag.Int("port", 3000, "Port to listen on")
	flag.Parse()
	courses.DownloadCourses()
	go Schedule()
	banner()
	server.Start(*portPtr)
}

func Schedule() {
	gocron.Every(1).Hour().Do(courses.DownloadCourses)
	<-gocron.Start()
}

func banner() {
	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("Quest For Excellence Learning Platform"))
	pterm.Info.Println("(c)2022 by Akhil Datla")
}