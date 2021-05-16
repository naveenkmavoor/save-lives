package controllers

import (
	"fmt"

	"encoding/json"
	"log"
	"os"
	"time"

	"../models"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/mongo"
)

type MqttController struct {
	session *mongo.Session
}

func NewMqttController(s *mongo.Session) *MqttController {
	return &MqttController{s}
}

func (mc MqttController) GetSensorVal() {

	var flag bool = false
	alertmssg := models.AlertMssg{}

	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("mqtt://tailor.cloudmqtt.com:12189")
	opts.SetUsername("frfnnxss")
	opts.SetPassword("RL27X-zpeWvS")
	opts.SetClientID("rasp-pi-go")
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println("SetDefaultPublishHandler : ", msg.Topic(), string(msg.Payload()))
	})

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe("some_topic", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//define a function for the default message handler
	var callback MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
		err := json.Unmarshal([]byte(msg.Payload()), &alertmssg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nAfter converting to object : %#v\n", alertmssg)

	}
	//subscribe
	if token := c.Subscribe("mqtt", 0, callback); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	for flag == false {
		time.Sleep(1 * time.Second)
	}

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("some_topic"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}
