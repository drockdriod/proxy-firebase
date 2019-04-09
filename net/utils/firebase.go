package utils

import (
	"encoding/json"

	"sync"

	"fmt"
  	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"

  	"golang.org/x/net/context"

  	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
  // "firebase.google.com/go/auth"

  	"google.golang.org/api/option"
)

var client *firestore.Client
var mqttClient mqtt.Client


func Connect(credentialFileLocation string, ctx context.Context) *firestore.Client {
	opt := option.WithCredentialsFile(credentialFileLocation)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
	  	fmt.Errorf("error initializing app: %v", err)
		
	}

	client, err = app.Firestore(ctx)
	if err != nil {
	  log.Fatalln(err)
	}
	// defer client.Close()

	return client
}

func GetRealtimeItemsByCollection(query *firestore.CollectionRef, waitGroup sync.WaitGroup, mClient mqtt.Client, ctx context.Context) {
	iter := query.Snapshots(ctx)

	defer iter.Stop()
	defer waitGroup.Done()


	for {
        qsnap, err := iter.Next()
        
        if err != nil {
            log.Println("err")
            log.Println(err)
        }
        fmt.Printf("At %s there were %d results.\n", qsnap.ReadTime, qsnap.Size)

		var items = []map[string]interface{}{}
        changes := qsnap.Changes

        for _, element := range changes {
        	if element.Kind == 2 {
        		items = append(items, element.Doc.Data())
        	}
		}

		data, err := json.Marshal(items)
		if err != nil {
			fmt.Println(err.Error())
		}

		mClient.Publish(query.ID, 0, false, string(data))




    }
}