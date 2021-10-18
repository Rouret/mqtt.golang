package mqtt

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe(topic string, qos byte, callback mqtt.MessageHandler) {

	if token := Global.MQTTClient.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
