package entity

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var(
	pickUp        =	"pick"
	createChannel =	"create"
)

func shift(s []string) (first string,slice []string){
	if len(s) == 0{
		return "",s
	}
	f := s[0]
	s = s[1:]
	return f,s
}

func callBot(s *discordgo.Session, m *discordgo.MessageCreate, commands []string){
	first,commands := shift(commands)
	if strings.HasPrefix(first,pickUp){
		pick(s,m,commands)
	}
}

func pick(s *discordgo.Session, m *discordgo.MessageCreate, options []string){
	picker := EquipmentPicker{"src/resources/"}
	survivor := NewSurvivor("unknown")
	survivor = picker.PickAllRandom(survivor)
	first,_ := shift(options)
	if strings.HasPrefix(first, fmt.Sprintf("%s %s", pickUp, "build")) {
		survivor = picker.PickBuildRandom(survivor)
	}
	SendPrivateMessage(s, m, survivor.Equipment())
}
