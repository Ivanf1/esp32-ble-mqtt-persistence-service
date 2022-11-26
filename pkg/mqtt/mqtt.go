package mqtt

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Ivanf1/esp32-mqtt-ble-persistence-service/pkg/db"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttMessagePublishHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

	i, err := strconv.Atoi(string(msg.Payload()))
	if err != nil {
		log.Printf("could not converto mqtt message payload to int")
		return
	}

	go db.Insert(i)
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
