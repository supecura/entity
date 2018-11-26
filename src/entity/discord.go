package entity

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/unicode/norm"
	"log"
	"strings"
	"time"
)

var (
	guild     *discordgo.Guild
	member    *discordgo.Member
	session   *discordgo.Channel
	st        []*discordgo.Role
	overwrite []*discordgo.PermissionOverwrite
)

func SendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)
	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}

func DeleteChannel(s *discordgo.Session, m *discordgo.MessageCreate) {

}

func (e Entity) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	picker := EquipmentPicker{"src/resources/"}
	fmt.Printf("%20s %20s %20s > %s\n", m.Author.ID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}
	var line = strings.Replace(string(norm.NFKC.Bytes([]byte(m.Content))), "  ", " ", -1)
	println(line)
	switch {
		case strings.HasPrefix(line, fmt.Sprintf("%s %s", e.BotName, pickUp)):
			survivor := NewSurvivor("unknown")
			survivor = picker.PickAllRandom(survivor)
			if strings.HasPrefix(line, fmt.Sprintf("%s %s %s", e.BotName, pickUp, "build")) {
				survivor = picker.PickBuildRandom(survivor)
			}
			SendMessage(s, c, survivor.Equipment())
			fmt.Printf("%20s %20s %20s > %s\n", m.Author.ID, time.Now().Format(time.Stamp), m.Author.Username, line)
	}
}
