package main

import (
	"fmt"

	"github.com/sandro-lex-symphony/gobot"
)

func Hello(sid, data string) {
	err := gobot.SendMessage(sid, "<card accent='tempo-bg-color--blue' iconSrc=''><body><h1>filler</h1></body></card>")
	gobot.Fatal(err)
}

func main() {
	gobot.Init()
	fmt.Println("XXXXX")
	gobot.Loop(Hello)
}
