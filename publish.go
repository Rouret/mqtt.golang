package mqtt

import (
	"encoding/json"
	"log"
)

func Send(topic string,qos int,message interface{},needToBeJson bool) {
	if(IsPersisten()){
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
	}else{
		log.Fatal("Need to be in PERSISTEN mode see LibConfiguration and Setup method")
	}
}