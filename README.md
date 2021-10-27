# MQTT PKG 


#### Install
`go get github.com/Rouret/mqtt.golang`

#### Documentation 

##### Setup

```go
package main

import "github.com/Rouret/mqtt.golang"

func main() {
	mqtt.Setup(mqtt.LibConfiguration{
		IsPersistent: true,
		BrokerUrl: "tcp://localhost",
    	BrokerPort: 1883,
		ID: 999,
	})
}
```

###### LibConfiguration

| Field | Type  | Description          
| :--------------- |:---------------:| :---------------:|
| IsPersistent  | boolean | That retains mqtt client value in memory between calls to the function          
          
##### Methods


| Method name | Input  | Output | Description | Need to be persistent          
| :--------------- |:---------------:| :---------------:| :---------------:| :---------------:|
| Setup  | setup LibConfiguration | void | Setup | 
| Connect  | void | mqtt.Client | Connect will create a connection to the message broker | 
| Send  | topic string, qos int, message interface{}, needToBeJson bool | void | [Publish will publish a message with the specified QoS and content to the specified topic.](https://github.com/eclipse/paho.mqtt.golang/blob/master/client.go) | X 
| Subscribe  | topic string, qos byte, callback mqtt.MessageHandler | void | [Subscribe starts a new subscription. Provide a MessageHandler to be executed when a message is published on the topic provided, or nil for the default handler.](https://github.com/eclipse/paho.mqtt.golang/blob/master/client.go) | X  

##### Example

```go
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
	mqtt.Connect("tcp://localhost:1883","123")
	mqtt.Send("temp",1,"message",false)

	//NO PERSITEN
	client := mqtt.Connect("tcp://localhost:1883","123")
	client.Publish("temp",1,false,"message")
	client.Connect().Wait()
}
```
