package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetOne(collectionName string, filter interface{}) (interface{}, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(HOST))
	if err != nil {
		return nil, err
	}

	db := client.Database(DB)
	collection := db.Collection(collectionName)
	if collection == nil {
		if err = db.CreateCollection(context.TODO(), collectionName); err != nil {
			return nil, err
		}
	}

	var result interface{}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
