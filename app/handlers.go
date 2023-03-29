package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestSample struct {
	Name   string `json:"name" bson:"name"`
	Header string `json:"header" bson:"header"`
	Body   string `json:"body" bson:"body"`
	Method string `json:"method" bson:"method"`
	Url    string `json:"url" bson:"url"`
}

type Testing struct {
	Title string `json:"title" bson:"title"`
}

// var testSamples = []TestSample{
// 	{
// 		Name:    "Example API Call",
// 		Headers: "Hello Jacob and Paul",
// 		Body:    "",
// 		Method:  "",
// 		Url:     ""},
// }

func getTest(c *gin.Context) {
	var test Testing
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

	// codenameDatabase := client.Database("codeName")
	testCollection := client.Database("codeName").Collection("tests")
	fmt.Println(testCollection)
	// filter := bson.D{{"title", "testing"}}

	err = testCollection.FindOne(context.TODO(), bson.D{}).Decode(&test)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, test)
}

func postTest(c *gin.Context) {
	var newTest TestSample

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
	testCollection := client.Database("codeName").Collection("tests")

	if err := c.BindJSON(&newTest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Empty Request Body")
		return
	}

	// Add the new album to the slice.
	insertResult, err := testCollection.InsertOne(ctx, newTest)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
}

func updateTest(c *gin.Context) {
	var newTest TestSample
	// var updatedTest TestSample

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

	if err := c.BindJSON(&newTest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Empty Request Body")
		return
	}

	testCollection := client.Database("codeName").Collection("tests")
	filter := bson.D{{Key: "name", Value: newTest.Name}}
	update := bson.D{{"$set", bson.D{
		{Key: "name", Value: newTest.Name},
		{Key: "header", Value: newTest.Header},
		{Key: "body", Value: newTest.Body},
		{Key: "method", Value: newTest.Method},
		{Key: "url", Value: newTest.Url},
	}}}

	result, err := testCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}

	fmt.Println(result.MatchedCount, result.ModifiedCount)
	// c.IndentedJSON(http.StatusOK, updatedTest)
}

func deleteTest(c *gin.Context) {
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

	name := c.Query("name")
	if name == "" {
		c.IndentedJSON(http.StatusBadRequest, "Query Parameter not valid")
		return
	}

	testCollection := client.Database("codeName").Collection("tests")
	filter := bson.D{{Key: "name", Value: name}}

	result, err := testCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, result.DeletedCount)
}
