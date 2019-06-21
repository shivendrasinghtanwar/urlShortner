package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"io/ioutil"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client = connectToMongo()
var shortUrlsCollection = client.Database("BillteTest").Collection("shortUrls")
var startHash = "AAAAAA"
func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/url/shorten",shortenUrlHandler)

	log.Printf("listening on port 5000")
    err := http.ListenAndServe(":5000", mux)
    if err != nil{
    	panic(err)
	}


}

func connectToMongo() *mongo.Client{
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




func shortenUrlHandler(w http.ResponseWriter, req *http.Request) {
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

		var res = resBody{
			ShortUrl:generateShortUrl(body.LongUrl),
		}

		jsonRes,objErr := json.Marshal(res)
		if objErr != nil {
			http.Error(w, "Marshalling Failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func generateShortUrl(longUrl string) string {

	var verr = client.Ping(context.TODO(), nil)
	if verr != nil {
		log.Fatal(verr)
		return "no"
	}

	var hash = genereateHash()
	var query = mongoQuery{
		hash,
	}
	cur, err1 := shortUrlsCollection.Find(context.TODO(), query)
	if err1 != nil {
		log.Fatal(err1.Error())
		return "no"
	}

	var results []*Record
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Record
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	fmt.Println(len(results))

	for len(results)>0{
		fmt.Println("here")
		hash = genereateHash()
		cur, err1 := shortUrlsCollection.Find(context.TODO(), query)
		if err1 != nil {
			log.Fatal(err1.Error())
			return "no"
		}
		for cur.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			var elem Record
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}

			results = append(results, &elem)
		}

		cur.Close(context.TODO())
	}

	return hash
}

func genereateHash() string {
	/*file, _ := ioutil.ReadFile("test.json")

	data := CatlogNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.CatlogNodes); i++ {
		fmt.Println("Product Id: ", data.CatlogNodes[i].Product_id)
		fmt.Println("Quantity: ", data.CatlogNodes[i].Quantity)
	}*/

	return ""
}