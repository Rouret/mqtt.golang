package main

import "github.com/Rouret/mqtt.golang"

func main() {

	mqtt.Setup(mqtt.LibConfiguration{
		IsPersistent: true,
		BrokerUrl: "tcp://localhost",
    	BrokerPort: 1883,
		ID: 999,
	})


	//PERSITEN
	mqtt.Connect()
	mqtt.Send("temp",1,"message",false)

	//NO PERSITEN
	// client := mqtt.Connect()
	// client.Publish("temp",1,false,"message")
	// client.Connect().Wait()
}