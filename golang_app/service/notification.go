package service

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/redis/go-redis/v9"
)

type NotificationService struct {
	app    *firebase.App
	client *messaging.Client
}

func NewNotificationService() (*NotificationService, error) {
	app := GetFirebaseApp()
	if app == nil {
		return nil, fmt.Errorf("firebase app non initialisée")
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Erreur lors de la récupération du client de notifications: %v", err)
		return nil, err
	}

	return &NotificationService{
		app:    app,
		client: client,
	}, nil
}

func (ns *NotificationService) SendPushNotification(token, title, body string) error {
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	_, err := ns.client.Send(context.Background(), message)
	if err != nil {
		return fmt.Errorf("échec de l'envoi de la notification: %w", err)
	}

	return nil
}

func (ns *NotificationService) StoreTokenInRedis(ctx context.Context, rdb *redis.Client, key, value string, expiration ...int) error {
	exp := 60
	if len(expiration) > 0 {
		exp = expiration[0]
	}
	key = fmt.Sprintf("%s:token", key)
	err := rdb.Set(ctx, key, value, time.Duration(exp)*time.Minute).Err()

	if err != nil {
		return err
	}
	return nil
}

func (ns *NotificationService) GetTokenFromRedis(ctx context.Context, rdb *redis.Client, key string) (string, error) {
	key = fmt.Sprintf("%s:token", key)
	token, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}
