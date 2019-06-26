package Structs

type Record struct {
	LongUrl string `json:"longUrl"`
	HashGen string `json:"hash"`
}

type BroadenUrlReqBody struct {
	HashGen string `json:"shortUrl"`
}

type BroadenUrlResBody struct {
	LongUrl string `json:"longUrl"`
}


type ReqBody struct {
	LongUrl string `json:"longUrl"`
}


type ResBody struct {
	ShortUrl string `json:"shortUrl"`
}

type CharMap struct{
	CharMap []charNode `json:"charMap"`
}

type charNode struct{
	Character string `json:"character"`
	Number string `json:"number"`
}

type Config struct{
	ShorturlHandlerPath string `json:"shortUrlHandlerPath"`
	BillerFrontEndPath string `json:"billerFrontendPath"`
	MonolithPath string `json:"monotlithPath"`
	PayerFrontEndPath string `json:"payerFrontendPath"`
	Mongo mongoConfig `json:"mongo"`
}

type mongoConfig struct{
	Database string `json:"database"`
	Collection string `json:"collection"`
	MongoPath string `json:"mongo_path"`
}