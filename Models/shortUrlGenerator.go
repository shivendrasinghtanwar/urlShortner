package Models

import (
	"context"
	"fmt"
	"log"
	"sync"
	"urlShortner/Database"
	. "urlShortner/Encoder"
	"urlShortner/Incrementor"
)

var mutex sync.Mutex
func GenerateShortUrl(longUrl string) string {

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
		return GenereateHashstringFromNumber(first.HashGen)
	}else{
		maxElem := allRecords[0]

		for _, value := range allRecords{
			if value.HashGen>maxElem.HashGen{
				maxElem = value
			}
		}

		maxElem.HashGen = Incrementor.AddOne(maxElem)
		maxElem.LongUrl = longUrl


		mutex.Lock()
		maxElem = Database.InsertHash(maxElem)

		mutex.Unlock()

		return GenereateHashstringFromNumber(maxElem.HashGen)
	}
}

