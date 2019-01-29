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

func (e Entity)shift(s []string) (first string,slice []string){
	if len(s) == 0{
		return "",s
	}
	f := s[0]
	s = s[1:]
	return f,s
}

func (e *Entity) callBot(s *discordgo.Session, m *discordgo.MessageCreate, commands []string){
	first,commands := e.shift(commands)
	if strings.HasPrefix(first,pickUp){
		e.pick(s,m,commands)
	}
}

func (e *Entity) pick(s *discordgo.Session, m *discordgo.MessageCreate, args []string){
	first,_ := e.shift(args)
	picker := EquipmentPicker{"src/resources/"}
	var role Role
	player := NewPlayer("unknown",role.Value("Survivor"))
	if strings.HasPrefix(first, fmt.Sprintf("%s", "killer")) {
		r := role.Value("killer")
		player.Role = r
	}
	player = picker.PickAllRandom(player)
	e.SendPrivateMessage(s, m, player.Equipment())
}
