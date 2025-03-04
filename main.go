package main

import (
	"mygo/controller/bot"
)

func main() {
	bot.Start()
	<-make(chan struct{})
}
