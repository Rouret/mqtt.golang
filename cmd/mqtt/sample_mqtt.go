package main

import "github.com/Rouret/mqtt.golang"

func main() {
	mqtt.Connect("tcp://localhost:1883","123")
	mqtt.Send("temp",1,"message",false)
}