package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongo() (*Mongo, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(HOST))
	if err != nil {
		return nil, err
	}
	return &Mongo{
		client: client,
		db:     client.Database(DB),
	}, nil
}

func (m *Mongo) GetOne(collectionName string, filter interface{}) (interface{}, error) {
	collection := m.db.Collection(collectionName)
	if collection == nil {
		return nil, errors.New("collection not found")
	}

	var result interface{}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *Mongo) GetAll(collectionName string, filter interface{}) ([]interface{}, error) {
	collection := m.db.Collection(collectionName)
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

func (m *Mongo) InsertOne(collectionName string, data interface{}) error {
	collection := m.db.Collection(collectionName)
	if collection == nil {
		return errors.New("collection not found")
	}

	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) UpdateOne(collectionName string, filter interface{}, update interface{}) error {
	collection := m.db.Collection(collectionName)
	if collection == nil {
		return errors.New("collection not found")
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) DeleteOne(collectionName string, filter interface{}) error {
	collection := m.db.Collection(collectionName)
	if collection == nil {
		return errors.New("collection not found")
	}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) Close() error {
	return m.client.Disconnect(context.TODO())
}
