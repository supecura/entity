package entity

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//ReadPark ... JSONファイル読み込み
func ReadPark(filePath string) []Park {
	bytes := readFile(filePath)
	var park []Park
	if err := json.Unmarshal(bytes, &park); err != nil {
		log.Fatal(err)
	}
	return park
}

//ReadPark ... JSONファイル読み込み
func ReadBuild(filePath string) []Build {
	bytes := readFile(filePath)
	var build []Build
	if err := json.Unmarshal(bytes, &build); err != nil {
		log.Fatal(err)
	}
	return build
}

//ReadItems ... JSONファイル読み込み
func ReadItems(filePath string) []Item {
	bytes := readFile(filePath)
	var item []Item
	if err := json.Unmarshal(bytes, &item); err != nil {
		log.Fatal(err)
	}
	return item
}

//ReadOffering ... JSONファイル読み込み
func ReadOffering(filePath string) []Offering {
	bytes := readFile(filePath)
	var offering []Offering
	if err := json.Unmarshal(bytes, &offering); err != nil {
		log.Fatal(err)
	}
	return offering
}

func readFile(filePath string) []byte{
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}