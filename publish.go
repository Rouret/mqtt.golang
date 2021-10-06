package mqtt

import (
	"encoding/json"
	"log"
)

func Send(topic string,qos int,message interface{},needToBeJson bool) {
	if(needToBeJson){
		jsonMessage, err := json.Marshal(message)
		if(err != nil){
			log.Fatal(err)
		}else{
			message = jsonMessage
		}
	}
	token := Global.MQTTClient.Publish(topic,byte(qos),false,message)
	token.Wait()
}