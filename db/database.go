package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	client      *mongo.Client
	databaseMap map[string]*mongo.Database
)

func Init(connectionUrl string) {
	var err error
	if client, err = mongo.Connect(context.TODO(), options.Client().
		ApplyURI(connectionUrl)); err != nil {
		log.Fatal(err)
	}

	databaseMap = map[string]*mongo.Database{}

	log.Println("Database initialized")
}

func GetDatabase(name string) *mongo.Database {
	if db, ok := databaseMap[name]; ok {
		return db
	}

	databaseMap[name] = client.Database(name)
	return databaseMap[name]
}
