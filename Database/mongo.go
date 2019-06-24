package Database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	. "urlShortner/Structs"
)

var client = ConnectToDB()
var shortUrlsCollection = client.Database("BillteTest").Collection("shortUrls")
func ConnectToDB() *mongo.Client{
	// Set client options

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	err = client.Connect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Connected to MongoDB!")

	return client
}

func InsertFirstHash(longUrl string) Record {

	first := Record{
		LongUrl:longUrl,
		HashGen:"010101010101",
	}
	_,err := shortUrlsCollection.InsertOne(context.TODO(),first)
	if err != nil{
		fmt.Println(err)
	}

	return first
}

func InsertHash(record Record) {
	_,err := shortUrlsCollection.InsertOne(context.TODO(),record)
	if err != nil{
		fmt.Println(err)
	}
}
/*
func findAll() (mongo.Cursor,error) {
	return shortUrlsCollection.Find(context.TODO(), bson.D{{}})
}*/