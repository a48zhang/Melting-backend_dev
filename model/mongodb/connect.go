package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var uri = "mongodb://localhost:27017/"

var Client *mongo.Client

func ConnectMongo() {
	tmp := os.Getenv("MONGO_URL")
	if tmp != "" {
		uri = tmp
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	var err error
	Client, err = mongo.Connect(context.TODO(), opts)
	var result bson.D
	err = Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	if err != nil {
		fmt.Println("connecting mongodb uri: ", uri)
		log.Fatal(err)
	}
}
