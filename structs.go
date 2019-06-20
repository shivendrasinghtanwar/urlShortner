package main

type Record struct {
	LongUrl string
	HashGen string
}

type reqBody struct {
	LongUrl string `json:"longUrl"`
}
