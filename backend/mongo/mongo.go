package mongo

import (
	"context"
	"errors"

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

func GetAll(collectionName string, filter interface{}) ([]interface{}, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(HOST))
	if err != nil {
		return nil, err
	}

	db := client.Database(DB)
	collection := db.Collection(collectionName)
	if collection == nil {
		return nil, errors.New("collection not found")
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var results []interface{}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func InsertOne(collectionName string, data interface{}) error {
	client, err := mongo.Connect(options.Client().ApplyURI(HOST))
	if err != nil {
		return err
	}

	db := client.Database(DB)
	collection := db.Collection(collectionName)
	if collection == nil {
		if err = db.CreateCollection(context.TODO(), collectionName); err != nil {
			return err
		}
	}

	_, err = collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOne(collectionName string, filter interface{}, update interface{}) error {
	client, err := mongo.Connect(options.Client().ApplyURI(HOST))
	if err != nil {
		return err
	}

	db := client.Database(DB)
	collection := db.Collection(collectionName)
	if collection == nil {
		return errors.New("collection not found")
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOne(collectionName string, filter interface{}) error {
	client, err := mongo.Connect(options.Client().ApplyURI(HOST))
	if err != nil {
		return err
	}

	db := client.Database(DB)
	collection := db.Collection(collectionName)
	if collection == nil {
		return errors.New("collection not found")
	}

	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
