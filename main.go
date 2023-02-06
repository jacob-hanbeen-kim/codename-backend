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

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	app.Start()
}

// func close(client *mongo.Client, ctx context.Context,
// 	cancel context.CancelFunc) {

// 	defer cancel()

// 	defer func() {
// 		if err := client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()
// }

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.
// func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(),
// 		30*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	fmt.Println("Connecting")
// 	return client, ctx, cancel, err
// }

// func ping(client *mongo.Client, ctx context.Context) error {

// 	// mongo.Client has Ping to ping mongoDB, deadline of
// 	// the Ping method will be determined by cxt
// 	// Ping method return error if any occurred, then
// 	// the error can be handled.
// 	if err := client.Ping(ctx, readpref.Primary()); err != nil {
// 		return err
// 	}
// 	fmt.Println("connected successfully")
// 	return nil
// }
