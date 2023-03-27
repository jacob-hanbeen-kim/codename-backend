package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jacob-hanbeen-kim/codename-backend/app/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb+srv://codename:codename123@cluster0.o3azdm6.mongodb.net/?retryWrites=true&w=majority
// testing

func main() {
	// Create a new client and connect to the server
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://codename:codename123@cluster0.o3azdm6.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	codenameDatabase := client.Database("codeName")
	testCollection := codenameDatabase.Collection("tests")
	// executionCollection := codenameDatabase.Collection("execution")
	// userCollection := codenameDatabase.Collection("user")

	testResponse, err := testCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "testing"},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(testResponse.InsertedID)
	app.Start()
}
