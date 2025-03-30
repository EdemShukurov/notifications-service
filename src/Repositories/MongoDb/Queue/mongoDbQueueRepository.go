package mongodb

import (
	"context"
	"errors"
	"fmt"
	queue "notifications-service/Repositories/Contracts/Queue"
	data "notifications-service/Repositories/Contracts/Queue/Data"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbNotificationRequestsQueue struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) queue.INotificationRequestsQueue {
	return &MongoDbNotificationRequestsQueue{
		collection: collection,
	}
}

func (r *MongoDbNotificationRequestsQueue) Enqueue(ctx context.Context, dbo data.NotificationRequestQueueItemDbo) error {
	_, err := r.collection.InsertOne(ctx, dbo)
	if err != nil {
		fmt.Errorf("Unable to insert doc in mongo. Error: %w", err)
		return err
	}

	fmt.Println("Doc is inserted in mongo, %d", dbo)

	return nil
}

func (r *MongoDbNotificationRequestsQueue) Peek(ctx context.Context, t string) (*data.NotificationRequestQueueItemDbo, error) {
	utcNow := primitive.NewDateTimeFromTime(time.Now().UTC())

	filter := bson.D{
		{"Type", t},
		{"Status", "pending"},
		{"SendAt", bson.M{"$lte": utcNow}}}

	update := bson.D{{"$set", bson.D{{"Status", "processing"}}}}

	opts := &options.FindOneAndUpdateOptions{
		Sort: bson.M{"SendAt": 1},
	}

	res := r.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, res.Err()
		// if !errors.Is(res.Err(), mongo.ErrNoDocuments) {
		// 	log.Error(ctx).Err(result.Err()).
		// 		Str("resource", c.name).
		// 		Msg("failed to get object from database")
		// }
	}

	elem := &data.NotificationRequestQueueItemDbo{}
	err := res.Decode(elem)
	if err != nil {
		return nil, err
	}

	return elem, nil
}

func (r *MongoDbNotificationRequestsQueue) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.D{{"_id", id}}, &options.DeleteOptions{})

	return err
}
