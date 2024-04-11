// utils/firebase.go
package utils

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	AuthClient      *auth.Client
	FirestoreClient *firestore.Client
)

func InitializeFirebaseClients() {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	AuthClient, err = firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	FirestoreClient, err = firebaseApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}
}
