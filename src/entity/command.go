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

func pick(s *discordgo.Session, m *discordgo.MessageCreate, args []string){
	first,_ := shift(args)
	picker := EquipmentPicker{"src/resources/"}
	var role Role
	player := NewPlayer("unknown",role.Value("Survivor"))
	if strings.HasPrefix(first, fmt.Sprintf("%s", "killer")) {
		r := role.Value("killer")
		player.Role = r
	}
	player = picker.PickAllRandom(player)
	SendPrivateMessage(s, m, player.Equipment())
}
