package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"urlShortner/Database"
	. "urlShortner/Models"
	. "urlShortner/Structs"
)




func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/url/shorten",shortenUrlHandler)
	mux.HandleFunc("/r/",broadenUrlHandler)
	mux.Handle("/",http.FileServer(http.Dir(Database.Configuration.StaticHtmlDirPath)))

	log.Printf("listening on port 5000")
    err := http.ListenAndServe(":5000", mux)
    if err != nil{
    	panic(err)
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

		if resErr != nil {
			http.Error(w, "Decoder Failed", http.StatusInternalServerError)
			return
		}

		if body.LongUrl=="" {
			http.Error(w, "Wrong URL to parse", http.StatusInternalServerError)
			return
		}

		//fmt.Println(body.LongUrl)

		res := ResBody{
			ShortUrl: GenerateShortUrl(body.LongUrl),
		}


		res.ShortUrl = Database.Configuration.ShorturlHandlerPath+res.ShortUrl

		fmt.Print("Short Url -")
		fmt.Println(res.ShortUrl)
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
	if req.Method == "GET" {

		shortUrl := strings.Split(req.URL.Path, "/")[2]

		fmt.Print("Got short Url ---")
		fmt.Println(shortUrl)

		if shortUrl=="" {
			http.Error(w, "Wrong URL to parse", http.StatusInternalServerError)
			return
		}

		res := BroadenUrlResBody{
			LongUrl: GenerateLongUrl(shortUrl),
		}

		http.Redirect(w, req, res.LongUrl,http.StatusMovedPermanently)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func serverIndexHandler(w http.ResponseWriter,req *http.Request) {
	if req.Method == "GET" {
		fmt.Println("got here")

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}









