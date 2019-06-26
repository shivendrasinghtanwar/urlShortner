package Encoder

import (
	"strings"
	. "urlShortner/CharMapReader"
)

var jsonData = ReadJsonCharMap()

func GenereateHashstringFromNumber(req string) string {

	/**
	For converting number string like '01010101' to 'AAAA'
	 */
	charArr := strings.Split(req,"")
	var res = ""

	for j:=len(charArr)-1;j>=0;j=j-2{
		var strToConv = charArr[j-1]+charArr[j]
		res = encode(strToConv) + res
	}

	return res
}


func encode(req string) string{
	//Converting from single digit like '01' to Characters like 'A'

	var res = ""
	for i := 0; i < len(jsonData.CharMap); i++ {
		if jsonData.CharMap[i].Number==req{
			res = jsonData.CharMap[i].Character
		}
	}
	return res
}