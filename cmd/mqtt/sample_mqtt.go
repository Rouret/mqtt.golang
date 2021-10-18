package main

import "github.com/Rouret/mqtt.golang"

func main() {

	mqtt.Setup(mqtt.LibConfiguration{
		IsPersistent: true,
	})

	//PERSITEN
	mqtt.Connect("tcp://localhost:1883","123")
	mqtt.Send("temp",1,"message",false)

	//NO PERSITEN
	client := mqtt.Connect("tcp://localhost:1883","123")
	client.Publish("temp",1,false,"message")
	client.Connect().Wait()
}