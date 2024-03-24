package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func (db *Database) Connect(uri string) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	db.Client = client

	return nil
}

func (db *Database) Disconnect() {
	db.Client.Disconnect(context.Background())
}

func (db *Database) Collection(name string) *mongo.Collection {
	return db.Client.Database("testing").Collection(name)
}
