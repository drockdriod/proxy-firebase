package utils

import (
	"encoding/json"

	"fmt"
  	"log"

	"github.com/drockdriod/proxy-firebase/mqttbroker"

  	"golang.org/x/net/context"

  	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
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

func GetRealtimeItemsByCollection(collection string) {
	query := client.Collection(collection)
	iter := query.Snapshots(ctx)
	defer iter.Stop()


	for {
        qsnap, err := iter.Next()
        // if err == auth.iterator.Done {
        //     break
        // }
        
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

		mqttbroker.Publish("test", string(data))




    }
}