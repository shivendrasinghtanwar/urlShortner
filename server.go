package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"urlShortner/Database"
	. "urlShortner/Structs"
)

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
	var table = Database.FindAll()
	i := len(table)
	for j:=0;j<i;j++{
		fmt.Println(genereateHash(table[j].HashGen))
	}
}

func shortenUrlHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		if req.Body == nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
			return
		}


		body := ReqBody{}
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

		res := ResBody{
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

	var verr = Database.Client.Ping(context.TODO(), nil)
	if verr != nil {
		log.Fatal(verr)
		return "no"
	}

	var allRecords = Database.FindAll()
	totalRecords := len(allRecords)
	if totalRecords==0 {
		fmt.Println("first Insert")
		first := Database.InsertFirstHash(longUrl)
		return genereateHash(first.HashGen)
	}else{
		maxElem := allRecords[0]

		for _, value := range allRecords{
			if value.HashGen>maxElem.HashGen{
				maxElem = value
			}
		}

		fmt.Println(maxElem.HashGen)

		intMaxElm,ierr1 := strconv.ParseInt(maxElem.HashGen, 10, 64)
		if ierr1 != nil {
			fmt.Println(ierr1)
		}

		incIntMaxElm := intMaxElm+1

		maxElem.HashGen = strconv.FormatInt(int64(incIntMaxElm),10)

		maxElem.HashGen = "0" + maxElem.HashGen
		maxElem.LongUrl = longUrl
		Database.InsertHash(maxElem)

		return genereateHash(maxElem.HashGen)
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

func readJsonCharMap() CharMap {
	file, err := ioutil.ReadFile("charMap.json")
	if err != nil {
		log.Fatal(err.Error())
		return CharMap{}
	}
	data := CharMap{}

	unmErr := json.Unmarshal([]byte(file), &data)
	if unmErr != nil {
		fmt.Println(unmErr)
		return CharMap{}
	}

	return data
}



