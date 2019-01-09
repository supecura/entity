package entity

import "github.com/bwmarrin/discordgo"

//SurvivorPlayer ...
type SurvivorPlayer struct {
	ID          string
	Park        []Park
	Item        Item
	addon       Addon
	Offering    Offering
	PickPattern PickPattern
}

func (player SurvivorPlayer) ShowPlayerInfo(s *discordgo.Session,c *discordgo.Channel){
	var message string
	message += "ID"+"\t"+ player.ID
	//SendMessage(s,c,message)
}

func (player SurvivorPlayer) GetPickPattern() PickPattern {
	return player.PickPattern
}

func (player SurvivorPlayer) SetPickPattern(p PickPattern) {
	player.PickPattern = p
}

//NewSurvivor
func NewSurvivor(id string) SurvivorPlayer {
	var p PickPattern
	survivor := SurvivorPlayer{ID: id,PickPattern:p.Value(id)}
	return survivor
}

func (player SurvivorPlayer) Equipment() string {
	var equipment string
	equipment += "Park"
	for i := 0; i < 4; i++ {
		equipment += "\n"+ player.Park[i].JapaneseName+"\t"+ player.Park[i].EnglishName
	}
	equipment += "\n"
	equipment += "\n"
	equipment += "Item"
	equipment += "\n"+ *player.Item.JapaneseName+"\t"+ *player.Item.EnglishName
	equipment += "\n"
	equipment += "\n"
	equipment += "Offering"
	equipment += "\n"+ *player.Offering.JapaneseName+"\t"+ *player.Offering.EnglishName
	return equipment
}

