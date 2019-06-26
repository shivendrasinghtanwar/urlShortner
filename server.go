package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	. "urlShortner/Models"
	. "urlShortner/Structs"
)




func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/url/shorten",shortenUrlHandler)
	mux.HandleFunc("/url/broaden",broadenUrlHandler)
	//mux.HandleFunc("/url/test",testHandler)


	log.Printf("listening on port 5000")
    err := http.ListenAndServe(":5000", mux)
    if err != nil{
    	panic(err)
	}


}

/*func testHandler(w http.ResponseWriter, req *http.Request) {

	elem := Record{
		HashGen:"010126262626",
		LongUrl:"SomeUrl",
	}

	fmt.Println(Incrementor.AddOne(&elem))

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}*/

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
			ShortUrl: GenerateShortUrl(body.LongUrl),
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

func broadenUrlHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		if req.Body == nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
			return
		}


		body := BroadenUrlReqBody{}
		decoder := json.NewDecoder(req.Body)

		resErr := decoder.Decode(&body)
		if resErr != nil {
			http.Error(w, "Decoder Failed", http.StatusInternalServerError)
			return
		}

		if body.HashGen=="" {
			http.Error(w, "Wrong URL to parse", http.StatusInternalServerError)
			return
		}

		res := BroadenUrlResBody{
			LongUrl: GenerateLongUrl(body.HashGen),
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









