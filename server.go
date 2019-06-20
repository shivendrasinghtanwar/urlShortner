package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/url/shorten",shortenUrl)


	log.Printf("listening on port 5000")
    err := http.ListenAndServe(":5000", mux)
    if err != nil{
    	panic(err)
	}

    connectToMongo()
}

func connectToMongo(){
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("test").Collection("trainers")

}

type reqBody struct {
	LongUrl string `json:"longUrl"`
}


func shortenUrl(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		if req.Body == nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
			return
		}


		body := reqBody{}
		decoder := json.NewDecoder(req.Body)

		resErr := decoder.Decode(&body)
		//resErr:= json.Unmarshal(req.Body,&body)
		if resErr != nil {
			http.Error(w, "Decoder Failed", http.StatusInternalServerError)
			return
		}

		if body.LongUrl=="" {
			http.Error(w, "Wrong URL to parse", http.StatusInternalServerError)
			return
		}

		getShortUrl(body.LongUrl)

		fmt.Println(body.LongUrl)
		res,objErr := json.Marshal(body)
		if objErr != nil {
			http.Error(w, "Marshalling Failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getShortUrl(longUrl string) string {
	//collection := client.Database("test").Collection("trainers")
	return ""
}