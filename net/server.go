package net

import (
	firebaseUtils "github.com/drockdriod/proxy-firebase/net/utils"
	"github.com/drockdriod/proxy-firebase/mqttbroker"
)

func Start(){

	mqttbroker.ConnectBroker()

	firebaseUtils.Connect()
	firebaseUtils.GetRealtimeItemsByCollection("messages")
}	