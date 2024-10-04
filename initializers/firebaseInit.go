package initializers

import (
	"context"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func FirebaseInitialization() {
	opt := option.WithCredentialsFile("/Users/sanda/Documents/Licenta/ExploreRO-server/authapp-39d6c-firebase-adminsdk-a7r8f-03c09bb114.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to connect to initialize firebase %v", err)
	}
	FirebaseApp = app
}
