package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//Global variable
var Global struct {
	MQTTClient mqtt.Client
	LibConfig LibConfiguration
} 
type LibConfiguration struct {
	IsPersistent bool
	BrokerUrl string
	BrokerPort int 
	ID int 
}

func Setup(setup LibConfiguration){
	Global.LibConfig = setup
}


func GetDefaultConfig() LibConfiguration{
	return LibConfiguration{
		IsPersistent: true,
		BrokerUrl: "tcp://localhost",
		BrokerPort: 1883,
		ID: 123,
	}
}


