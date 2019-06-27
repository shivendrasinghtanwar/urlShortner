package Incrementor

import (
	"fmt"
	"strconv"
	"strings"
	. "urlShortner/Structs"
)

var totalCharLength = int64(27)

func AddOne(elem *Record) string {
	hashCharArr := strings.Split(elem.HashGen,"")
	for i:= len(hashCharArr)-1;i >= 0;i=i-2{
		parseDigitString := hashCharArr[i-1] + hashCharArr[i]

		parseDigit,err := strconv.ParseInt(parseDigitString, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		if parseDigit == totalCharLength-1{
			fmt.Println("Add Next one!")
			hashCharArr = addOneInNextParseDigit(hashCharArr,i)
			break
		}

		if parseDigit < totalCharLength{
			incrementedParseDigit := parseDigit+1
			if incrementedParseDigit < totalCharLength{
				if incrementedParseDigit < 10 {
					hashCharArr[i-1] = "0"
					hashCharArr[i] = strconv.FormatInt(incrementedParseDigit,10)
				} else {
					incrementedParseString := strconv.FormatInt(incrementedParseDigit,10)
					incrementedParseCharArr := strings.Split(incrementedParseString,"")
					hashCharArr[i-1] = incrementedParseCharArr[0]
					hashCharArr[i] = incrementedParseCharArr[1]
				}
				break
			}
		}
	}

	return strings.Join(hashCharArr,"")
}

func addOneInNextParseDigit(hashCharArr []string,i int) []string{
	ones := hashCharArr[i-2]
	tens := hashCharArr[i-3]
	parseDigitString := tens + ones

	parseDigit,err := strconv.ParseInt(parseDigitString, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Print("parse digit--->>>")
	//fmt.Println(parseDigit)

	if parseDigit == totalCharLength-1{
		//fmt.Println("here")
		hashCharArr = addOneInNextParseDigit(hashCharArr,i-2)
		return hashCharArr
	}

	if parseDigit < totalCharLength{
		incrementedParseDigit := parseDigit+1
		if incrementedParseDigit < totalCharLength{
			if incrementedParseDigit < 10 {
				tens = "0"
				ones = strconv.FormatInt(incrementedParseDigit,10)
			} else {
				incrementedParseString := strconv.FormatInt(incrementedParseDigit,10)
				incrementedParseCharArr := strings.Split(incrementedParseString,"")
				tens = incrementedParseCharArr[0]
				ones = incrementedParseCharArr[1]
			}

			for j:=i;j< len(hashCharArr);j=j+2{
				hashCharArr[j-1]="0"
				hashCharArr[j]="1"
			}
			hashCharArr[i-2] = ones
			hashCharArr[i-3] = tens
		}
	}
	return hashCharArr
}