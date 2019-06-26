package FileReader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	. "urlShortner/Structs"
)

func ReadJsonCharMap() CharMap {
	file, err := ioutil.ReadFile("charMap.json")
	if err != nil {
		log.Fatal(err.Error())
		return CharMap{}
	}
	data := CharMap{}

	unmErr := json.Unmarshal([]byte(file), &data)
	if unmErr != nil {
		fmt.Println(unmErr)
		return CharMap{}
	}

	return data
}

func ReadConfig() Config {
	data := Config{}
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err.Error())
		return data
	}


	unmErr := json.Unmarshal([]byte(file), &data)
	if unmErr != nil {
		fmt.Println(unmErr)
		return data
	}

	return data

}
