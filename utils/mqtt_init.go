package utils

import (
	"crypto/tls"
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
)

var batchSize = 100
var deviceBuffer []AllDevice

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var device AllDevice

	err := json.Unmarshal(msg.Payload(), &device)
	if err != nil {
		log.Printf("Error parsing message: %v", err)
		return
	}

	deviceBuffer = append(deviceBuffer, device)

	if len(deviceBuffer) >= batchSize {
		insertData()
	}
}

func MqttClientInit() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("MQTT_BROKER_USERNAME")
	password := os.Getenv("MQTT_BROKER_PASSWORD")
	clientID := os.Getenv("MQTT_BROKER_CLIENTID")
	protocol := os.Getenv("MQTT_BROKER_PROTOCOL")
	host := os.Getenv("MQTT_BROKER_HOST")
	port := os.Getenv("MQTT_BROKER_PORT")

	uri := protocol + host + ":" + port
	opts := mqtt.NewClientOptions().AddBroker(uri)

	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	opts.SetTLSConfig(tlsConfig)

	opts.OnConnect = func(c mqtt.Client) {
		log.Println("Connected to MQTT broker")
	}
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Printf("Connection lost: %v\n", err)
	}

	var mqttClient = mqtt.NewClient(opts)

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	topics := []string{"library/reference", "library/information_technology", "library/medical","library/filipiniana","library/publication","library/serials",}

	for _, topic := range topics {
		token := mqttClient.Subscribe(topic, 2, messageHandler)
		token.Wait()

		if token.Error() != nil {
			log.Printf("Failed to subscribe to topic %s: %v", topic, token.Error())
		} else {
			log.Printf("Subscribe to topic: %s\n", topic)
		}
	}

	for {
		time.Sleep(1 * time.Second)
	}
}
