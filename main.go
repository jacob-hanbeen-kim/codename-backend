package main

import (
	"context"

	"github.com/jacob-hanbeen-kim/codename-backend/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

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

// // This is a user defined method that returns mongo.Client,
// // context.Context, context.CancelFunc and error.
// // mongo.Client will be used for further database operation.
// // context.Context will be used set deadlines for process.
// // context.CancelFunc will be used to cancel context and
// // resource associated with it.
// func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(),
// 		30*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
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

// require (
// 	github.com/gin-contrib/sse v0.1.0 // indirect
// 	github.com/gin-gonic/gin v1.8.1 // indirect
// 	github.com/go-playground/locales v0.14.0 // indirect
// 	github.com/go-playground/universal-translator v0.18.0 // indirect
// 	github.com/go-playground/validator/v10 v10.11.1 // indirect
// 	github.com/goccy/go-json v0.9.11 // indirect
// 	github.com/golang/snappy v0.0.1 // indirect
// 	github.com/json-iterator/go v1.1.12 // indirect
// 	github.com/klauspost/compress v1.13.6 // indirect
// 	github.com/leodido/go-urn v1.2.1 // indirect
// 	github.com/mattn/go-isatty v0.0.16 // indirect
// 	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
// 	github.com/modern-go/reflect2 v1.0.2 // indirect
// 	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
// 	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
// 	github.com/pkg/errors v0.9.1 // indirect
// 	github.com/ugorji/go/codec v1.2.7 // indirect
// 	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
// 	github.com/xdg-go/scram v1.1.1 // indirect
// 	github.com/xdg-go/stringprep v1.0.3 // indirect
// 	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
// 	go.mongodb.org/mongo-driver v1.11.1 // indirect
// 	golang.org/x/crypto v0.1.0 // indirect
// 	golang.org/x/net v0.1.0 // indirect
// 	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
// 	golang.org/x/sys v0.1.0 // indirect
// 	golang.org/x/text v0.4.0 // indirect
// 	google.golang.org/protobuf v1.28.1 // indirect
// 	gopkg.in/yaml.v2 v2.4.0 // indirect
// )
