package mqtt

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe(topic string, qos byte, callback mqtt.MessageHandler) {
	if IsPersisten() {
		if token := Global.MQTTClient.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	} else {
		log.Fatal("Need to be in PERSISTEN mode see LibConfiguration and Setup method")
	}
}
