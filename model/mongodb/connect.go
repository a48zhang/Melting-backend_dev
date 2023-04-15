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
	uri = os.Getenv("MONGO_URL")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	var err error
	Client, err = mongo.Connect(context.TODO(), opts)
	err = Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&bson.D{})
	if err != nil {
		fmt.Println("connecting mongodb uri: ", uri)
		log.Fatal(err)
	}
}
