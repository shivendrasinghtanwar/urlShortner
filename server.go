package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	//"io/ioutil"

	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var client = connectToMongo()
var shortUrlsCollection = client.Database("BillteTest").Collection("shortUrls")

func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/url/shorten",shortenUrlHandler)
	mux.HandleFunc("/url/test",testHandler)


	log.Printf("listening on port 5000")
    err := http.ListenAndServe(":5000", mux)
    if err != nil{
    	panic(err)
	}


}

func testHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("here")
	//genereateHash()
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

		fmt.Println(body.LongUrl)

		res := resBody{
			ShortUrl:generateShortUrl(body.LongUrl),
		}

		jsonRes,objErr := json.Marshal(res)
		if objErr != nil {
			http.Error(w, "Marshalling Failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_,rerr := w.Write(jsonRes)
		if rerr !=nil {
			fmt.Println(rerr)
		}


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


	cur,err1 := shortUrlsCollection.Find(context.TODO(), bson.D{{}})
	if err1 != nil {
		fmt.Println(err1)
		return ""
	}else {
		var scur = cur
		if scur.Next(context.TODO()) {
			var results []*Record
			var maxElm Record
			var elem Record

			err := cur.Decode(&maxElm)
			if err != nil {
				log.Fatal(err)
			}
			for cur.Next(context.TODO()) {

				fmt.Println("got here2")
				err := cur.Decode(&elem)
				if err != nil {
					log.Fatal(err)
				}
				intMaxElm,ierr1 := strconv.ParseInt(maxElm.HashGen,10,64)
				if ierr1 != nil {
					fmt.Println(ierr1)
				}

				intElem,ierr2 := strconv.ParseInt(elem.HashGen,10,64)
				if ierr2 != nil {
					fmt.Println(ierr2)
				}

				fmt.Println("got here")
				fmt.Print(intMaxElm)
				fmt.Print("   ")
				fmt.Println(intElem)

				if intMaxElm < intElem {
					maxElm = elem
				}

				results = append(results, &elem)
			}

			intMaxElm,ierr1 := strconv.ParseInt(maxElm.HashGen, 10, 64)
			if ierr1 != nil {
				fmt.Println(ierr1)
			}

			incIntMaxElm := intMaxElm+1

			maxElm.HashGen = strconv.FormatInt(int64(incIntMaxElm),10)

			maxElm.HashGen = "0" + maxElm.HashGen
			maxElm.LongUrl = longUrl
			insertHash(maxElm)

			cur.Close(context.TODO())
			return genereateHash(maxElm.HashGen)
		} else {
			fmt.Println("first Insert")
			first := insertFirstHash(longUrl)
			return genereateHash(first.HashGen)
		}
	}
}

func genereateHash(req string) string {

	charArr := strings.Split(req,"")
	var res = ""

	for j:=len(charArr)-1;j>=0;j=j-2{
		var strToConv = charArr[j-1]+charArr[j]
		res = encode(strToConv) + res
	}

	return res
}

func encode(req string) string{
	//Converting from numbers to Characters
	data := readJsonCharMap()

	var res = ""
	for i := 0; i < len(data.CharMap); i++ {
		if data.CharMap[i].Number==req{
			res = data.CharMap[i].Character
		}
	}
	return res
}

func readJsonCharMap() charMap{
	file, err := ioutil.ReadFile("charMap.json")
	if err != nil {
		log.Fatal(err.Error())
		return charMap{}
	}
	data := charMap{}

	unmErr := json.Unmarshal([]byte(file), &data)
	if unmErr != nil {
		fmt.Println(unmErr)
		return charMap{}
	}

	return data
}

func insertFirstHash(longUrl string) Record {

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

func insertHash(record Record) {
	_,err := shortUrlsCollection.InsertOne(context.TODO(),record)
	if err != nil{
		fmt.Println(err)
	}
}