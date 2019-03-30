package mqttbroker
// courtesy of: https://www.cloudmqtt.com/docs/go.html

import (
	"fmt"
	"log"
	// "net/url"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientId string, uri string) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri))
	opts.SetClientID(clientId)
	return opts
}

func listen(uri string, topic string) {
	client := connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func Publish(topic string, payload string) {
	clientName := os.Getenv("MQTT_CLIENT")
	port := os.Getenv("MQTT_PORT")
	
	uri := fmt.Sprintf("%s:%s",clientName,port)
	client := connect("pub", uri)

	client.Publish(topic, 0, false, payload)

}

func ConnectBroker() {
	clientName := os.Getenv("MQTT_CLIENT")
	port := os.Getenv("MQTT_PORT")
	
	uri := fmt.Sprintf("%s:%s",clientName,port)


	topic := "test"

	go listen(uri, topic)

	client := connect("pub", uri)
		client.Publish(topic, 0, false, "hello there")
	// timer := time.NewTicker(1 * time.Second)
	// for t := range timer.C {
	// }
}