package main

import (
  "fmt"
  "log"

  "golang.org/x/net/context"

  firebase "firebase.google.com/go"
  // "firebase.google.com/go/auth"

  "google.golang.org/api/option"
)

func main(){
	ctx := context.Background()

	opt := option.WithCredentialsFile("./firebase/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
	  	fmt.Errorf("error initializing app: %v", err)
		
	}

	client, err := app.Firestore(ctx)
	if err != nil {
	  log.Fatalln(err)
	}
	defer client.Close()

	fmt.Println("All messages:")
	iter := client.Collection("messages").Documents(ctx)
	defer iter.Stop()

	for {
        doc, err := iter.Next()
        // if err == auth.iterator.Done {
        //     break
        // }
        
        if err != nil {
            log.Fatalln(err)
        }
        if doc.Exists() {
        	fmt.Println(doc.Data())
        }
	}
}