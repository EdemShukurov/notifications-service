package services

import (
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddMongoDbService(mongoConnectionString string) (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	//serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	//opts := options.Client().ApplyURI(mongoConnectionString).SetServerAPIOptions(serverAPI)
	opts := options.Client().ApplyURI(mongoConnectionString)
	fmt.Println("Mongo client options TYPE: ", reflect.TypeOf(opts), "\n")

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		fmt.Errorf("Unable to connect to Mongo. Error: %w", err)
		//panic(err)
		return nil, err
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		//panic(err)
	// 	}
	// }()

	return client, nil
}
