package main

type Record struct {
	LongUrl string `json:"longUrl"`
	HashGen string `json:"hashGenerated"`
}

type reqBody struct {
	LongUrl string `json:"longUrl"`
}


type resBody struct {
	ShortUrl string `json:"shortUrl"`
}

type mongoQuery struct {
	HashGen string `json:"hashGen"`
}