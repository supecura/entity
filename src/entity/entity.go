package entity

import "github.com/bwmarrin/discordgo"

func New(botName string) (entity Entity,e error){
	entity = Entity{BotName:botName}
	return entity , e
}

func (e Entity) AddDiscordClient(client *discordgo.Session) {
	e.DiscordClient = client
}