package Structs

type Record struct {
	LongUrl string `json:"longUrl"`
	HashGen string `json:"hash"`
}

type ReqBody struct {
	LongUrl string `json:"longUrl"`
}


type ResBody struct {
	ShortUrl string `json:"shortUrl"`
}

type mongoQuery struct {
	HashGen string `json:"hashGen"`
}

type CharMap struct{
	CharMap []charNode `json:charMap`
}

type charNode struct{
	Character string `json:"character"`
	Number string `json:"number"`
}