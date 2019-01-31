package entity

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//ReadPark ... JSONファイル読み込み
func (picker *EquipmentPicker) ReadPark(filePath string) []Park {
	bytes := picker.readFile(filePath)
	var park []Park
	if err := json.Unmarshal(bytes, &park); err != nil {
		log.Fatal(err)
	}
	return park
}

//ReadPark ... JSONファイル読み込み
func (picker *EquipmentPicker) ReadBuild(filePath string) []Build {
	bytes := picker.readFile(filePath)
	var build []Build
	if err := json.Unmarshal(bytes, &build); err != nil {
		log.Fatal(err)
	}
	return build
}

//ReadItems ... JSONファイル読み込み
func (picker *EquipmentPicker) ReadItems(filePath string) []Item {
	bytes := picker.readFile(filePath)
	var item []Item
	if err := json.Unmarshal(bytes, &item); err != nil {
		log.Fatal(err)
	}
	return item
}

//ReadOffering ... JSONファイル読み込み
func (picker *EquipmentPicker) ReadOffering(filePath string) []Offering {
	bytes := picker.readFile(filePath)
	var offering []Offering
	if err := json.Unmarshal(bytes, &offering); err != nil {
		log.Fatal(err)
	}
	return offering
}

func (picker *EquipmentPicker) readFile(filePath string) []byte{
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}