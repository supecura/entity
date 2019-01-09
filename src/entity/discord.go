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

func (e Entity) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("%20s %20s %20s > %s\n", m.Author.ID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
	var line = strings.Replace(string(norm.NFKC.Bytes([]byte(m.Content))), "  ", " ", -1)
	sp := strings.Split(line, " ")
	first,commands := shift(sp)
	if strings.HasPrefix(first,e.BotName){
		callBot(s,m,commands)
	}
	fmt.Printf("%20s %20s %20s > %s\n", m.Author.ID, time.Now().Format(time.Stamp), m.Author.Username, line)
}

func SendPrivateMessage(discordSession *discordgo.Session, m *discordgo.MessageCreate, message string) {
	userChannel,err := discordSession.UserChannelCreate(m.Author.ID)
	if err != nil {
		log.Println("Error create channel: ", err)
	}
	discordSession.ChannelMessageSend(userChannel.ID, message)
}