package test

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"./proxyfire"
)

func test(){
	clientName := os.Getenv("MQTT_CLIENT")
	port := os.Getenv("MQTT_PORT")

	pc := proxyfire.Connect("./firebase/serviceAccount.json", clientName, port)

	pc.RunListeners()
}