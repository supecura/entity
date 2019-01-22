package entity

//Park ...
type Park struct {
	JapaneseName string `json:"japaneseName"`
	EnglishName  string `json:"englishName"`
	Type         string `json:"type"`
}

//Offering ...
type Offering struct {
	JapaneseName *string `json:"japaneseName"`
	EnglishName  *string `json:"englishName"`
	Type         string  `json:"type"`
	Rarity       *string `json:"rarity"`
}

//addon ...
type Addon struct {
	JapaneseName string `json:"japaneseName"`
	EnglishName  string `json:"englishName"`
	Type         string `json:"type"`
}

//Item ...
type Item struct {
	JapaneseName *string	`json:"japaneseName"`
	EnglishName  *string	`json:"englishName"`
	Type         string 	`json:"type"`
	Rarity       *string	`json:"rarity"`
}

//Build ...
type Build struct {
	Park		 	[]Park		`json:"park"`
	Description		string		`json:"description"`
	Item        	Item		`json:"item"`
	Offering		Offering	`json:"offering"`
}

type Entity struct{
}

