package mqtt

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttMessagePublishHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var client mqtt.Client

func ClientSetup() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(os.Getenv("MQTT_HOST"))
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	opts.SetDefaultPublishHandler(mqttMessagePublishHandler)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func SubscribeAndListen() {
	if token := client.Subscribe(os.Getenv("MQTT_SUBSCRIBE_TOPIC"), 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
