package utils

import (
	"encoding/json"

	"fmt"
  	"log"

	"github.com/drockdriod/proxy-firebase/mqttbroker"

  	"golang.org/x/net/context"

  	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"sync"
  // "firebase.google.com/go/auth"

  	"google.golang.org/api/option"
)

var client *firestore.Client
var ctx context.Context


func Connect() *firestore.Client {
	ctx = context.Background()

	opt := option.WithCredentialsFile("./firebase/serviceAccount.json")
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

func GetRealtimeItemsByAllCollections() {
	iter := client.Collections(ctx)

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

		go GetRealtimeItemsByCollection(query, waitGroup)
    }
}

func GetRealtimeItemsByCollection(query *firestore.CollectionRef, waitGroup sync.WaitGroup) {
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
		log.Println(query.ID)
		log.Println(items)

		data, err := json.Marshal(items)
		if err != nil {
			fmt.Println(err.Error())
		}


		mqttbroker.Publish(query.ID, string(data))




    }
}