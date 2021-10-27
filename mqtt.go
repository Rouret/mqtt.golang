package mqtt

import (
	"log"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)


func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

func Connect() mqtt.Client {

	log.Println("TIPS: See setup method to config the mqtt connexion")
	endPoint := Global.LibConfig.BrokerUrl+":"+strconv.Itoa(Global.LibConfig.BrokerPort)
	clientId := strconv.Itoa(Global.LibConfig.ID)

	log.Print("Trying to connect (" + endPoint + "," + clientId + ")...")

	opts := createClientOptions(endPoint, clientId )
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}else{
		if(Global.LibConfig.IsPersistent){
			Global.MQTTClient = client
		}
	}
	
	return client
}





