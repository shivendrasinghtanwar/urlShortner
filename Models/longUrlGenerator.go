package Models

import (
	"urlShortner/Database"
	. "urlShortner/Encoder"
)

func GenerateLongUrl(shortUrlHash string) string{

	decodedHash := GenerateNumberFromHashString(shortUrlHash)
	//fmt.Println(decodedHash)

	responseRec := Database.FindOneByHashNumber(decodedHash)
	//fmt.Print(responseRec)

	return responseRec.LongUrl
}
