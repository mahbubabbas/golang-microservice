package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DB_CONN_STRING = "mongodb://127.0.0.1:27017"
	DB_NAME = "schooldb"
	STUDENT_COLLECTION = "students"
	USER_COLLECTION = "users"
	PASS_KEY = "dsjsdjfjsdfsdjfsdfsdfjsdf$@skk%^sdkksdfkdsf732838"
)

var mClient *mongo.Client
var mDB *mongo.Database

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_CONN_STRING))

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		panic(err)
	}

	log.Println("DB connection is successful")

	mClient = client
	mDB = client.Database(DB_NAME)
}

func DisconnectMongo() {
	mClient.Disconnect(context.TODO())
}
