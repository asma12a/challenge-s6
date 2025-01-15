package service

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App

func InitializeFirebase() (*firebase.App, error) {
	credentials := ""

	opt := option.WithCredentialsJSON([]byte(credentials))

	var err error
	app, err = firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: "",
	}, opt)
	if err != nil {
		log.Fatalf("erreur lors de l'initialisation de l'application: %v", err)
	}

	return app, nil
}

func GetFirebaseApp() *firebase.App {
	InitializeFirebase()
	return app
}
