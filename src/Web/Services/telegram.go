package services

import (
	"fmt"
	service "notifications-service/Domain/Telegram"
	sender "notifications-service/Domain/Telegram/Infrastructure/Sending"
	repositories "notifications-service/Repositories/MongoDb/Notifications"
	queue "notifications-service/Repositories/MongoDb/Queue"
	controller "notifications-service/Web/Controllers/Telegram"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddTelegramService(telegramBotApiToken string, mongoClient *mongo.Client) controller.ITelegramController {
	senderInstance, err := sender.New(telegramBotApiToken)
	if err != nil {
		fmt.Errorf("Unable to initialize Telegram sender. Error: %w", err)
		return nil
	}

	// "mongodb://localhost:27017/"
	coll := mongoClient.Database("globi-notifications-service").Collection("queue")

	q := queue.New(coll)
	repo := repositories.New(coll)

	serviceInstance := service.New(senderInstance, q, repo)
	return controller.New(serviceInstance)
}
