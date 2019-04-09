package models
import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"cloud.google.com/go/firestore"
	firebaseUtils "github.com/drockdriod/proxy-firebase/net/utils"
	"sync"
  	"golang.org/x/net/context"

  	"log"
)

type ProxyClient struct {
	Client *Client
	Ctx context.Context
}

type Client struct {
	FbClient *firestore.Client
	MqttClient mqtt.Client
}


func (pc *ProxyClient) RunListeners() {
	getRealtimeItemsByAllCollections(pc)
}

func getRealtimeItemsByAllCollections(pc *ProxyClient) {
	client := pc.Client.FbClient
	mqttClient := pc.Client.MqttClient
	iter := client.Collections(pc.Ctx)

	var waitGroup sync.WaitGroup
    var collections []string


	for {
		ref, err := iter.Next()

		if err != nil {
            log.Println("GetRealtimeItemsByAllCollections error")
            log.Println(err.Error())
            break
        } else {
			log.Println(ref.ID)
			collections = append(collections, ref.ID)
        	
        }


	}

	waitGroup.Add(len(collections))

    defer waitGroup.Wait()

    for _, collection := range collections {
		query := client.Collection(collection)

		go firebaseUtils.GetRealtimeItemsByCollection(query, waitGroup, mqttClient, pc.Ctx)
    }
}