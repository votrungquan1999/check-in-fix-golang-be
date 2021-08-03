package setup

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
	"log"
)

var AuthClient *auth.Client
var FirestoreClient *firestore.Client
var StorageClient *storage.Client

func StartFirebase() {
	config := GetFirebaseConfig()
	configJSON, err := json.Marshal(config)
	if err != nil {
		log.Fatalf("error marshall config: %v\n", err)
	}

	ctx := context.Background()
	//opt := option.WithCredentialsFile("/Users/quan.vo/Documents/checkin_fix/checkin-fix-firebase-adminsdk.json")
	opt := option.WithCredentialsJSON(configJSON)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	AuthClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing auth client: %v\n", err)
	}

	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing firestore client: %v\n", err)
	}

	StorageClient, err = app.Storage(ctx)
	if err != nil {
		log.Fatalf("error initializing storage client: %v\n", err)
	}

	//var err error
	//AuthClient, err = storage.NewClient(ctx, option.WithCredentialsFile("/Users/quan.vo/Documents/checkin_fix/checkin-fix-firebase-adminsdkabc.json"))
	//if err != nil {
	//	log.Fatalf("error initializing client: %v\n", err)
	//}
}
