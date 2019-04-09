package proxyfire

import (
	firebaseUtils "github.com/drockdriod/proxy-firebase/net/utils"
	"github.com/drockdriod/proxy-firebase/mqttbroker"
	// server "github.com/drockdriod/proxy-firebase/net"
	"github.com/drockdriod/proxy-firebase/net/models"
  	"golang.org/x/net/context"

)

func Connect(fbServiceAccountLocation string, mqttUri string, port int) (*models.ProxyClient) {
	ctx := context.Background()
	
	mqttClient := mqttbroker.ConnectBroker(mqttUri, port)
	fb := firebaseUtils.Connect(fbServiceAccountLocation, ctx)
	client := &models.ProxyClient{
		&models.Client{
			fb, 
			mqttClient,
		},
		ctx,
	}


	return client
}
