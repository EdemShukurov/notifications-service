package mongodb

import (
	"context"
	"fmt"
	contract "notifications-service/Repositories/Contracts/Notifications"
	data "notifications-service/Repositories/Contracts/Notifications/Data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbNotificationsRepository struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) contract.INotificationsRepository {
	return &MongoDbNotificationsRepository{
		collection: collection,
	}
}

func (r *MongoDbNotificationsRepository) CreateOrUpdate(ctx context.Context, dbo data.NotificationDbo) error {
	filter := bson.D{{"_id", dbo.Id}}
	opts := options.Replace().SetUpsert(true)

	// Updates a documents or inserts a document if no documents are
	// matched and prints the results
	_, err := r.collection.ReplaceOne(ctx, filter, dbo, opts)
	if err != nil {
		fmt.Errorf("Unable to insert doc in mongo. Error: %w", err)
		return err
	}

	fmt.Println("Doc is inserted in mongo, %d", dbo)

	return nil
}
