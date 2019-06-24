package main

type Record struct {
	LongUrl string `json:"longUrl"`
	HashGen string `json:"hash"`
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

type charMap struct{
	CharMap []charNode `json:charMap`
}

type charNode struct{
	Character string `json:"character"`
	Number string `json:"number"`
}