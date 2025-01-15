package service

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"github.com/asma12a/challenge-s6/config"

)

var app *firebase.App

func InitializeFirebase() (*firebase.App, error) {
	credentialsFile := ""

	opt := option.WithCredentialsFile(credentialsFile)
	var err error
	app, err = firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: config.Env.ProjectID,
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
