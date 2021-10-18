package mqtt

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//Global variable
var Global struct {
	MQTTClient mqtt.Client
	LibConfig LibConfiguration
} 
type LibConfiguration struct {
	IsPersistent bool
}

func Setup(setup LibConfiguration){
	Global.LibConfig = setup
}

func IsPersisten() bool{
	return Global.LibConfig.IsPersistent;
}

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

func Connect(brokerURI string, clientId string) mqtt.Client {
	log.Print("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}else{
		if(IsPersisten()){
			Global.MQTTClient = client
		}
	}
	
	return client
}
