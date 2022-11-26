package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ivanf1/esp32-mqtt-ble-persistence-service/pkg/db"
	"github.com/Ivanf1/esp32-mqtt-ble-persistence-service/pkg/mqtt"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	keepAlive := make(chan os.Signal, 1)
	signal.Notify(keepAlive, os.Interrupt, syscall.SIGTERM)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mqtt.ClientSetup()
	mqtt.SubscribeAndListen()

	db.Connect()

	<-keepAlive
}
