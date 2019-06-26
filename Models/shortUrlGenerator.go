package Models

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"urlShortner/Database"
	. "urlShortner/Encoder"

)

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

		return GenereateHashstringFromNumber(maxElem.HashGen)
	}
}
